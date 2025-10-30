#!/usr/bin/env bash
usage() { echo "Usage: $0 [-v <version>] [-b <branch-prefix>] [-j <jira-ticket>]" 1>&2; exit 1; }

VERSION=""
BRANCH_PREFIX=""
JIRA=""

while getopts v:b:j: arg; do
    case "${arg}" in
        v)
            VERSION="${OPTARG}"
            ;;
        b)
            BRANCH_PREFIX="${OPTARG}"
            ;;
        j)
            JIRA="${OPTARG}"
            ;;
        *)
            usage
            ;;
    esac
done

# Validate required arguments
if [ -z "${VERSION}" ] || [ -z "${BRANCH_PREFIX}" ] || [ -z "${JIRA}" ]; then
    usage
fi

echo "Version: $VERSION"
echo "Branch: $BRANCH_PREFIX"
echo "Jira: $JIRA"
