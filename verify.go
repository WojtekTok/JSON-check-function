package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type IAMRolePolicy struct {
	PolicyDocument PolicyDocument `json:"PolicyDocument"`
	PolicyName     string         `json:"PolicyName"`
}

type PolicyDocument struct {
	Version   string      `json:"Version"`
	Statement []Statement `json:"Statement"`
}

type Statement struct {
	Sid      string   `json:"Sid"`
	Effect   string   `json:"Effect"`
	Action   []string `json:"Action"`
	Resource string   `json:"Resource"`
}

// Function expected in a task:
func VerifyJSON(policy IAMRolePolicy) bool {
	for _, statement := range policy.PolicyDocument.Statement {
		if statement.Resource == "*" {
			return false
		}
	}
	return true
}

// Not sure if this validation is necessary, but otherwise it may not be the expectet data
func validatePolicy(policy IAMRolePolicy) error {
	if policy.PolicyName == "" {
		return fmt.Errorf("missing 'PolicyName")
	}

	if policy.PolicyDocument.Version == "" {
		return fmt.Errorf("missing 'PolicyDocument.Version")
	}

	if len(policy.PolicyDocument.Statement) == 0 {
		return fmt.Errorf("missing 'PolicyDocument.Statement")
	}

	for i, stm := range policy.PolicyDocument.Statement {
		if len(stm.Action) == 0 {
			return fmt.Errorf("missing 'PolicyDocument.Statement.Action' in statement: %d", i)
		}
		if stm.Effect == "" {
			return fmt.Errorf("missing 'PolicyDocument.Statement.Effect' in statement: %d", i)
		}
		if stm.Resource == "" {
			return fmt.Errorf("missing 'PolicyDocument.Statement.Resource' in statement: %d", i)
		}
	}
	return nil
}

func GetPath() string {
	fmt.Println("Path of AWS::IAM::Role Policy JSON file:")
	var path string
	fmt.Scanln(&path)
	return path
}

func ReadJSON(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("Error occured while reading file: %s\n", err)
	} else {
		var policy IAMRolePolicy
		err := json.Unmarshal(file, &policy)
		if err != nil {
			return fmt.Sprintf("Error occured while unmarshalling JSON file: %s\n", err)
		} else {
			isValid := validatePolicy(policy)
			if isValid == nil {
				return fmt.Sprintf("Doesn't contain single asterisk: %t", VerifyJSON(policy))
			} else {
				return fmt.Sprintf("JSON input data incorrect: %s", isValid)
			}

		}
	}
}

func main() {
	fmt.Println(ReadJSON(GetPath()))
}
