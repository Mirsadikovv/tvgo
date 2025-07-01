<script setup lang="ts">
<<- $root := .>>
<<- $serviceNameSc := toSnake $root.ServiceName >>
<<- $serviceNameUc := toCamel $root.ServiceName >>
<<- $serviceNameLc := toLowerCamel $root.ServiceName >>

import { ref } from "vue";
import { <<$serviceNameUc>>Service , type <<$serviceNameUc>>PageType , type <<$serviceNameUc>>PartialType } from "@/service";

import PerPage from "@/components/PerPage.vue";
import ColumsPiker from "@/components/ColumsPiker.vue";
import PageLoading from "@/components/PageLoading.vue";
import Table from "@/components/quasar/table/Table.vue";
import IconBtn from "@/components/quasar/btn/IconBtn.vue";
import Button from "@/components/quasar/btn/Button.vue";
import SearchAutocomplete from "@/components/quasar/search/SearchAutocomplete.vue";
import Search from "@/components/quasar/search/Search.vue";
import Form from "@/components/quasar/form/Form.vue";
import Expasion from "@/components/quasar/Expasion.vue";
import TablePaginate from "@/components/quasar/table/TablePaginate.vue";

import { useRouter } from "vue-router";
const router = useRouter();

const models = ref< <<$serviceNameUc>>PageType >({
	data: [],
	totalRows: 0,
	totalPages: 0,
	pageSize: 0,
	currentPage: 0,
});

const pick = {
	id: false,
	edit: false,
};

const pikers = ref({});

async function find(query: string) {
	const response = await <<$serviceNameUc>>Service.getAll<< $serviceNameUc >>(query)
	if (!response) return;
	models.value = response
}

let searchModel = ref< <<$serviceNameUc>>PartialType >({});

async function save() {
	searchModel.value = {};
	await find("");
	router.replace({ query: {} });
	return true;
}
</script>

<template>
	<PageLoading :find="find" #="{ loading, fetch }">
		<div class="flex! gap-x-4 items-center mb-3">
			<q-breadcrumbs>
				<q-breadcrumbs-el :label="$tl('<< toLower $serviceNameSc >>_item')" icon="article" />
				<q-breadcrumbs-el :label="$tl('page_for_table')" />
			</q-breadcrumbs>
			<q-space></q-space>
			<ColumsPiker
				:name="$tl('columns')"
				v-model="pikers"
				:picker="pick"
			/>
			<PerPage
				@per-page="fetch"
				:display-value="$tl('perpage')"
			></PerPage>

			<Button 
				v-if="$auth.canPage( `CREATE_<< toUpper $serviceNameSc >>` )" 
					:to="{
						name: `CREATE_<< toUpper $serviceNameSc >>`,
					}"
				>
				{{ $tl("create") }}
			</Button>
			
		</div>

		<Expasion>
			<Form :model-value="searchModel" :save="save">
				<template #fio="{ model }">
					<Search
						:label="$tl('fullName')"
						class="col-lg-3 col-md-6 col-sm-12 col-xs-12"
						v-model="model.fullname"
						name="fullname"
						@search="fetch"
						input-debounce="500"
					></Search>
				</template>

				<template #status="{ model }">
					<SearchAutocomplete
						class="col-lg-2 col-md-6 col-sm-12 col-xs-12"
						v-model="model.status"
						label="status"
						query-name="status"
						@update="fetch"
						:find="async () => ['ACTIVE', 'INACTIVE']"
						:option-label="(option) => $tl(option)"
					/>
				</template>

				<template #actions="{ loading }">
					<q-space></q-space>
					<div class="flex! gap-x-4 items-center">
						<q-btn
							@click="save()"
							:loading="loading"
							class="bg-white text-primary pt1! px-8!"
							outline
						>
							{{ $tl("cancel") }}
						</q-btn>
					</div>
				</template>
			</Form>
		</Expasion>

		<Table :loading="loading" :models="models" :pick="pikers" hasOrder>
			<template #name:thead> </template>
			<template #name></template>

			<template #edit:thead>
				<div class="text-center">
					{{ $tl("action") }}
				</div>
			</template>
			<template #edit="{ model }">
				<div class="text-center">
					<IconBtn 
						v-if="$can( `EDIT_<< toUpper $serviceNameSc >>` )"
						:key="model.id" icon="edit" 
						class="mr-2"
						:to="{
							name: `EDIT_<< toUpper $serviceNameSc >>`,
							params: { id: model.id },
						}"
					/>
					<IconBtn 
						v-if="$can( `VIEW_<< toUpper $serviceNameSc >>` )"
						:key="model.id" icon="visibility" 
						class="mr-2"
						:to="{
								name: `VIEW_<< toUpper $serviceNameSc >>`,
								params: { id: model.id },
							}"
					/>
				</div>
			</template>

			<template #tfoot="{ totalPages }">
				<TablePaginate :total="totalPages" @page="fetch" />
			</template>
		</Table>
	</PageLoading>
</template>
