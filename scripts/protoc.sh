#!/usr/bin/env bash

GO_MODULE="gitlab.com/naftis/app/naftis"

set -e
set -u
set -o pipefail

function generate_golang() {
  package="${1}"
  file="${2}"

  echo "Generating golang sources for ${file}@${package} ..."
  protoc --go_opt=module="${GO_MODULE}" --plugin="protoc-gen-go-grpc" --go_out=plugins=grpc:. "api/protoc/${package}/${file}.proto"
  protoc-go-inject-tag -input="pkg/protocol/${package}/${file}.pb.go"
}

function verify() {
  program="${1}"

  if ! command -v ${program} &> /dev/null
  then
      echo "ERROR: '${program}' binary could not be found"
      exit 1
  fi
}

function main() {
  echo "This script generates protobuf sources."

  verify "protoc"
  verify "protoc-gen-go-grpc"
  verify "protoc-go-inject-tag"

  generate_golang "blockchain" "node_label_message"
  generate_golang "blockchain" "contract_request_message"
  generate_golang "blockchain" "contract_response_message"
  generate_golang "blockchain" "workload_specification_message"

  generate_golang "api" "list_observed_workloads_request_message"
  generate_golang "api" "list_observed_workloads_response_message"
  generate_golang "api" "list_scheduled_workloads_request_message"
  generate_golang "api" "list_scheduled_workloads_response_message"
  generate_golang "api" "schedule_workload_request_message"
  generate_golang "api" "schedule_workload_response_message"
  generate_golang "api" "api_service"

  generate_golang "entity" "scheduled_workload_message"
  generate_golang "entity" "observed_workload_message"
  generate_golang "entity" "workload_spec_message"

  echo -e "Done!"
}

main "${@:-}"