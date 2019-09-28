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

substitute_values() {
    input="./templates/${1}.go.tmpl"
    while IFS= read -r line
    do
    KIND_LOWERCASE=$(echo $KIND | tr '[:upper:]' '[:lower:]')
    KIND_SHORT=$(echo $KIND | sed 's/MongoDBAtlas//g')
    sed 's/_KIND_LOWERCASE_/'${KIND_LOWERCASE}'/g' $line > kind_lower_replaced_line.tmp
    sed 's/_KIND_/'${KIND}'/g' kind_lower_replaced_line.tmp > kind_replaced_line.tmp
    sed 's/_KIND_SHORT_/'${KIND_SHORT}'/g' kind_replaced_line.tmp > kind_short_replaced_line.tmp
    sed 's/_API_VERSION_/'${API_VERSION}'/g' kind_short_replaced_line.tmp > substitued_values.go
    done < "$input"
    rm kind_lower_replaced_line.tmp kind_replaced_line.tmp kind_short_replaced_line.tmp
}



pushd "${SCRIPT_DIR}/../" > /dev/null
check_input_vars
# handle add_kind.gp.tmpl
substitute_values add_kind
mv substitued_values.go pkg/controller/add_${KIND_LOWERCASE}.go
# handle kind_controller.gp.tmpl
substitute_values kind_controller
mv substitued_values.go pkg/controller/${KIND_LOWERCASE}/${KIND_LOWERCASE}_controller.go
popd > /dev/null
