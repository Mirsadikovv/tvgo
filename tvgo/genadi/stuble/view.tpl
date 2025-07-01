<script setup lang="ts">
<<- $root := .>>
<<- $serviceNameSc := toSnake $root.ServiceName >>
<<- $serviceNameUc := toCamel $root.ServiceName >>
<<- $serviceNameLc := toLowerCamel $root.ServiceName >>
import { useRouter } from "vue-router";
import { <<$serviceNameUc>>Service , <<$serviceNameUc>>PartialType } from "@/services";
import { ref } from "vue";



export interface Props {
	id: number | string;
}
const { id } = defineProps<Props>();
const ID = $computed(() => +id);

const router = useRouter();

let model = ref< <<$serviceNameUc>>PartialType >({});

<<$serviceNameUc>>Service.get<< $serviceNameUc >>ById(ID).then((data) => {
 	model.value = data;
});
</script>

<template>
	<div class="flex! gap-x-4 items-center mb-6">
		<q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
		<q-breadcrumbs>
			<q-breadcrumbs-el :label="$tl('<< toLower $serviceNameSc >>_item')" icon="article" />
			<q-breadcrumbs-el :label="$tl('view_page')" />
		</q-breadcrumbs>
	</div>
	<q-markup-table separator="cell" flat bordered>
		<tbody>
			<tr>
				<td class="font-bold text-left w-1/2">{{ $tl("id") }}</td>
				<td>{{ model?.id }}</td>
			</tr>
		</tbody>
	</q-markup-table>
</template>
