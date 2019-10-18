#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "$0")" ; pwd -P)"
show_help() {
    echo "Usage: "
    echo
    echo "  ./`basename "$0"` --api-version v1alpha1 --kind KIND"
}
POSITIONAL=()
while [[ $# -gt 0 ]]; do
    key="$1"
    case $key in
        --api-version)
        API_VERSION="$2"
        shift
        shift
        ;;
        -k|--kind)
        KIND="$2"
        shift
        shift
        ;;
        *)
        show_help
        exit 1
        ;;
    esac
done

check_input_vars() {
    if [[ ${API_VERSION} == "" ]] || [[ ${KIND} == "" ]]; then
        show_help
        exit 1
    fi
}

check_if_already_substitued() {
    file=$1
    replace=y
    # handle add_kind.gp.tmpl
    if [[ -f ${file} ]]; then
        read -p "[warn] file ${file} already exists. Do you want to replace it? (Y/n) " replace
    else
        replace=y
    fi
    case $replace in
        [Yy]* ) echo "y";;
        [Nn]* ) echo "n";;
        * ) echo "n";;
    esac
}

substitute_values() {
    input="./templates/${1}.go.tmpl"
    while IFS= read -r line
    do
    sed 's/_KIND_LOWERCASE_/'${KIND_LOWERCASE}'/g' $line > kind_lower_replaced_line.tmp
    sed 's/_KIND_SHORT_/'${KIND_SHORT}'/g' kind_lower_replaced_line.tmp > kind_short_replaced_line.tmp
    sed 's/_KIND_/'${KIND}'/g' kind_short_replaced_line.tmp > kind_replaced_line.tmp
    sed 's/_API_VERSION_/'${API_VERSION}'/g' kind_replaced_line.tmp > substitued_values.go
    done < "$input"
    rm kind_lower_replaced_line.tmp kind_replaced_line.tmp kind_short_replaced_line.tmp
}



pushd "${SCRIPT_DIR}/../" > /dev/null
check_input_vars
KIND_LOWERCASE=$(echo $KIND | tr '[:upper:]' '[:lower:]')
KIND_SHORT=$(echo $KIND | sed 's/MongoDBAtlas//g')

# create controller dir
[[ ! -d pkg/controller/${KIND_LOWERCASE} ]] && mkdir pkg/controller/${KIND_LOWERCASE}

# handle add_kind.go.tmpl
return_add_kind=$(check_if_already_substitued pkg/controller/add_${KIND_LOWERCASE}.go)
if [[ ${return_add_kind} == "y" ]];then 
    substitute_values add_kind 
    mv substitued_values.go pkg/controller/add_${KIND_LOWERCASE}.go
fi
# handle kind_controller.go.tmpl
return_kind_controller=$(check_if_already_substitued pkg/controller/${KIND_LOWERCASE}/${KIND_LOWERCASE}_controller.go)
if [[ ${return_kind_controller} == "y" ]];then 
    substitute_values kind_controller
    mv substitued_values.go pkg/controller/${KIND_LOWERCASE}/${KIND_LOWERCASE}_controller.go
fi
# handle kind_controller_test.go.tmpl
return_kind_controller_test=$(check_if_already_substitued pkg/controller/${KIND_LOWERCASE}/${KIND_LOWERCASE}_controller_test.go)
if [[ ${return_kind_controller_test} == "y" ]];then 
    substitute_values kind_controller_test
    mv substitued_values.go pkg/controller/${KIND_LOWERCASE}/${KIND_LOWERCASE}_controller_test.go
fi
# handle e2e/kind_test.go.tmpl
return_kind_controller_test=$(check_if_already_substitued test/e2e/${KIND_LOWERCASE}_test.go)
if [[ ${return_kind_controller_test} == "y" ]];then 
    substitute_values e2e/kind_test
    mv substitued_values.go test/e2e/${KIND_LOWERCASE}_test.go
fi
popd > /dev/null
