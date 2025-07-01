<script setup lang="ts">
<<- $root := .>>
<<- $serviceNameSc := toSnake $root.ServiceName >>
<<- $serviceNameUc := toCamel $root.ServiceName >>
<<- $serviceNameLc := toLowerCamel $root.ServiceName >>


import { ref } from "vue";
import { useRouter } from "vue-router";
import { <<$serviceNameUc>>Service, type <<$serviceNameUc>>PartialType } from "@/service";
import { formRequired } from "@/common";

import Title from "@/components/title.vue";
import Form from "@/components/quasar/form/Form.vue";
import Input from "@/components/quasar/form/Input.vue";
import Button from "@/components/quasar/btn/Button.vue";
import Autocomplete from "@/components/quasar/form/Autocomplete.vue";

const router = useRouter();

const model = ref< <<$serviceNameUc>>PartialType >({});

async function save(newModel: <<$serviceNameUc>>PartialType ) {
    const response = await <<$serviceNameUc>>Service.create<< $serviceNameUc >>(newModel)

	if (!response) return false;

    router.push({
		name: `VIEW_<< toUpper $serviceNameSc >>`,
		params: { id: response.id },
	});

    return true;
}
</script>

<template>
    <div class="flex! gap-x-4 items-center mb-3">
        <q-btn flat color="accent" icon="arrow_back" @click="router.back()" />
        <q-breadcrumbs>
            <q-breadcrumbs-el :label="$tl('<<$serviceNameLc>>_item')" icon="article" :to="{ name: 'PAGE_<<toUpper $serviceNameSc>>' }"/>
            <q-breadcrumbs-el :label="$tl('page_for_create')" />
        </q-breadcrumbs>
    </div>

    <Form :model-value="model" :save="save">
        <template #title>
            <Title class="mb2">{{ $tl("add_<<$serviceNameLc>>") }} </Title>
        </template>

        <template #defaultInput="{ model }">
           <Input v-model="model.name" label="name" class="col-lg-4 col-md-6 col-12" :rules="[formRequired()]"/>
        </template>

        <template #defaultSelect="{ model }">
            <Autocomplete
				v-model="model.status"
				:rules="[formRequired()]"
				label="status"
				class="col-lg-4 col-md-6 col-12"
				:find="async () => ['ACTIVE', 'INACTIVE']"
				:option-label="(option) => $tl(option)"
			/>
        </template>

      <template #actions="{ loading }">
			<q-space></q-space>
			<Button :loading="loading" type="submit">
				{{ $tl("save") }}
			</Button>
		</template>
    </Form>
</template>