package purple

import "testing"

func TestParseART(t *testing.T) {
	art := `attack_technique: T1609
display_name: Kubernetes Exec Into Container
atomic_tests:
- name: ExecIntoContainer
  auto_generated_guid: d03bfcd3-ed87-49c8-8880-44bb772dea4b
  description: |
    Attackers who have permissions, can run malicious commands in containers in the cluster using exec command (“kubectl exec”). In this method, attackers can use legitimate images, such as an OS image (e.g., Ubuntu) as a backdoor container, and run their malicious code remotely by using “kubectl exec”.
  supported_platforms:
  - linux
  - macos
  input_arguments:
    namespace:
      description: K8s namespace to use
      type: String
      default: default
    command:
      description: Command to run
      type: String
      default: uname
  executor:
    prereq_command: |
      which kubectl
    command: |
      kubectl create -f src/busybox.yaml -n #{namespace}
      kubectl exec -n #{namespace} busybox -- #{command}
    cleanup_command: |
      kubectl delete pod busybox -n #{namespace}
    name: bash
    elevation_required: false`

	artResult, err := ParseART(art)
	if err != nil  {
		t.Fatal(err)
		t.Fail()
	}

	if artResult == nil {
		t.Fail()
	}

	if artResult.DisplayName != "Kubernetes Exec Into Container" {
		t.Fail()
	}
}
