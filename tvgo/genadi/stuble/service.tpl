<<- $root := .>>
<<- $serviceNameSc := toSnake $root.ServiceName >>
<<- $serviceNameUc := toCamel $root.ServiceName >>
<<- $serviceNameLc := toLowerCamel $root.ServiceName >>
<<- $modelName := printf "%sModel" $serviceNameUc >>

<<- $model := index $root.JsonData.Definitions $modelName >>

import { Try } from "@/common";
import { api } from "@/plugins/axios.plugin";
import type { PageDataType } from "@/service";

export type << $serviceNameUc >>Type = {
    <<- range $field, $property := $model.Properties >>
    << $field >>: << swaggerType $property.Type >>;
    <<- end >>
};

export type << $serviceNameUc >>PartialType = Partial< << $serviceNameUc >>Type >;
export type << $serviceNameUc >>PageType = PageDataType< << $serviceNameUc >>Type >;

<<- $usedModels := dict >> 

<<- range $path, $methods := $root.JsonData.Paths >>
    <<- if contains $path (toLower $serviceNameUc) >>
        <<- range $method, $methodValues := $methods >>
            <<- $params := $methodValues.Parameters >>
            <<- $responses := $methodValues.Responses >>

            <<- range $param, $paramValue := $params >>
                <<- if and (eq $paramValue.In "body") $paramValue.Schema $paramValue.Schema.Ref >>
                    <<- $dto := normalizeModelName (extractDtoName $paramValue.Schema.Ref) >>
                    <<- $_ := set $usedModels $dto true >>
                <<- end >>
            <<- end >>

            <<- range $status, $respValue := $responses >>
                <<- if and (hasPrefix $status "2") $respValue.Schema $respValue.Schema.Ref >>
                    <<- $dto := normalizeModelName (extractDtoName $respValue.Schema.Ref) >>
                    <<- $_ := set $usedModels $dto true >>
                <<- end >>
            <<- end >>
        <<- end >>
    <<- end >>
<<- end >>

<<- range $modelName, $_ := $usedModels >>
    <<- $model := index $root.JsonData.Definitions $modelName >>
    <<- if $model >>
export type << $modelName >> = {
        <<- range $field, $property := $model.Properties >>
    << $field >>: << swaggerType $property.Type >>;
        <<- end >>
};
export type << $modelName >>Partial = Partial< << $modelName >> >;

    <<- end >>
<<- end >>


class << $serviceNameUc >>Service {
    <<- range $path, $methods := $root.JsonData.Paths >>
    <<- if contains $path (toLower $serviceNameUc) >>
    <<- range $method, $methodValues := $methods >>
    <<- $params := $methodValues.Parameters >>
    <<- $responses := $methodValues.Responses >>
    <<- $hasBody := false >>
    <<- $hasPath := false >>
    <<- $hasQuery := false >>
    <<- $dto := "" >>

    @Try(
       {
        <<- range $param, $paramValue := $params >>
        <<- if eq $paramValue.In "body" >>

            async onSuccess(result) {
                (await import("@/common/Notify")).SuccesNotify(result.statusText);
            },

        <<- $hasBody = true >>
        <<- $dto =  $paramValue.Schema.Ref >>
        <<- end >>
        <<- end >>

            async onError(err) {
                (await import("@/common/Notify")).ErrorNotify(
                    // @ts-ignore
                    err?.response?.data.message || err.message,
                );
		    },
        }

    )
    async << toLowerCamel $methodValues.OperationID >>(
        <<- range $param, $paramValue := $params >>
        <<- if eq $paramValue.In "path" >>
        << $paramValue.Name >>: string,
        <<- else if eq $paramValue.In "query" >>
        // << $paramValue.Name >>?: string,
        <<- else if eq $paramValue.In "body" >>
        body: << $serviceNameUc >>PartialType,
        <<- $hasBody = true >>
        <<- $dto =  $paramValue.Schema.Ref >>
        <<- end >>
        <<- end >>
        query?: string,
    ) {
        <<- range $param, $paramValue := $params >>
        <<- if eq $paramValue.In "body" >>
        <<- $hasBody = true >>
        <<- $dto =  $paramValue.Schema.Ref >>
        <<- end >>
        <<- end >>
        <<- if $hasBody >>
        const { data } = await api.<< $method | toLower >>< << extractDtoName $dto >> >(
            `<< replacePath $path >>`,
            body
        );
        <<- else if eq $method "put" >>
        const { data } = await api.<< $method | toLower >>< << $serviceNameUc >>Type >(
            `<< replacePath $path >>`
        );
        <<- else if eq $method "get" >>
        const { data } = await api.<< $method | toLower >>< << $serviceNameUc >>PartialType >(
            `<< replacePath $path >>/${query}`
        );
        <<- else if eq $method "delete" >>
        const { data } = await api.<< $method | toLower >>(
            `<< replacePath $path >>`
        );
        <<- else >>
        const { data } = await api.<< $method | toLower >>< any >(
            `<< replacePath $path >>`
        );
        <<- end >>

        return data;
    }
    <<- end >>
    <<- end >>
    <<- end >>
}


const << $serviceNameLc >>Service = new << $serviceNameUc >>Service();
export { << $serviceNameLc >>Service as << $serviceNameUc >>Service };
