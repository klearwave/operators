export GIT_REPO_URL="github.com/klearwave/operators-capability"

# defines usage when both a capability and component are expected
capability_component_usage() {
    echo "Usage: $0 <capability> <component>"
    echo "  capability: the capability name (directory under capabilities/)"
    echo "  component:  the component name (the project name, e.g. cert-manager)"
}

# defines function to ensure capability is provided
ensure_capability_input() {
    : ${CAPABILITY?missing environment variable CAPABILITY}
    echo "CAPABILITY: ${CAPABILITY}"
}

# defines function to ensure component is provided
ensure_component_input() {
    : ${COMPONENT?missing environment variable COMPONENT}
    echo "COMPONENT: ${COMPONENT}"
}

# defines function to ensure both capability and component are provided
ensure_capability_component_input() {
    ensure_capability_input
    ensure_component_input
}

# defines function to set commmon variables for capability
set_capability_vars() {
    ensure_capability_input

    CAPABILITY_DIR_BASE=capabilities
    CAPABILITY_DIR=${CAPABILITY_DIR_BASE}/${CAPABILITY}
    CAPABILITY_DIR_SOURCE=${CAPABILITY_DIR}/.source

    echo "CAPABILITY_DIR: ${CAPABILITY_DIR}"
}

# defines function to set commmon variables for component
set_component_vars() {
    set_capability_vars
    ensure_component_input

    COMPONENT_DIR=${CAPABILITY_DIR_SOURCE}/${COMPONENT}

    echo "COMPONENT_DIR: ${COMPONENT_DIR}"
}

# defines function to set commmon variables for capability and component
set_capability_component_vars() {
    set_capability_vars
    set_component_vars
}
