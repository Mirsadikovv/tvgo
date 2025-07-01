import type { RouteRecordRaw } from "vue-router";


<<- $root := .>>
<<- $serviceNameUp := toUpper $root.ServiceName >>
<<- $serviceNameLo := toLower $root.ServiceName >>
<<- $serviceNameSc := toSnake $root.ServiceName >>
<<- $serviceNameUc := toCamel $root.ServiceName >>
<<- $serviceNameLc := toLowerCamel $root.ServiceName >>
<<- $serviceNameKe := toKebab $root.ServiceName >>

export const <<$serviceNameLc>>ListRoute: RouteRecordRaw = {
    path: "<<$serviceNameKe>>",
    name: "PAGE_<<toUpper $serviceNameSc>>",
    component: () => import("@/modules/ContentService/<<$serviceNameUc>>/pages/Page.vue"),
    meta: {
        title: "<<$serviceNameLc>>_item_page",
        activeLinkGroup: "<<$serviceNameUp>>_GROUP",
        sidebar: {
            icon: "api",
            label: "<<$serviceNameSc>>_item",
        },
    },
};

export const <<$serviceNameLc>>CreateRoute: RouteRecordRaw = {
    path: "create-<<$serviceNameKe>>",
    name: "CREATE_<<toUpper $serviceNameSc>>",
    component: () => import("@/modules/ContentService/<<$serviceNameUc>>/pages/Create.vue"),
    meta: {
        title: "<<$serviceNameLc>>_item_create",
        activeLinkGroup: "<<$serviceNameUp>>_GROUP",
    },
};

export const <<$serviceNameLc>>UpdateRoute: RouteRecordRaw = {
    path: "edit-<<$serviceNameKe>>/:id",
    name: "EDIT_<<toUpper $serviceNameSc>>",
    props: true,
    component: () => import("@/modules/ContentService/<<$serviceNameUc>>/pages/Edit.vue"),
    meta: {
        title: "<<$serviceNameLc>>_item_edit",
        activeLinkGroup: "<<$serviceNameUp>>_GROUP",
    },
};

export const <<$serviceNameLc>>ViewRoute: RouteRecordRaw = {
    path: "view-<<$serviceNameKe>>/:id",
    name: "VIEW_<<toUpper $serviceNameSc>>",
    props: true,
    component: () => import("@/modules/ContentService/<<$serviceNameUc>>/pages/View.vue"),
    meta: {
        title: "<<$serviceNameLc>>_item_view",
        activeLinkGroup: "<<$serviceNameUp>>_GROUP",
    },
};