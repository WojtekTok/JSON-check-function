package main

import (
	"strings"
	"testing"
)

func TestReadJSON(t *testing.T) {
	t.Run("Valid file", func(t *testing.T) {
		result := ReadJSON("example.json")
		if !strings.Contains(result, "Doesn't contain single asterisk") {
			t.Errorf("Expected message for correct verification, got: %s", result)
		}
	})

	t.Run("Invalid file path", func(t *testing.T) {
		result := ReadJSON("nonexistant.json")
		if !strings.Contains(result, "Error occured while reading file") {
			t.Errorf("Expected error message for nonexistent file, got: %s", result)
		}
	})
}

func TestVerify(t *testing.T) {
	tests := []struct {
		name   string
		policy IAMRolePolicy
		want   bool
	}{
		{
			name: "Single asterisk",
			policy: IAMRolePolicy{
				PolicyDocument: PolicyDocument{
					Statement: []Statement{{Resource: "*"}},
				},
			},
			want: false,
		},
		{
			name: "Double asterisk",
			policy: IAMRolePolicy{
				PolicyDocument: PolicyDocument{
					Statement: []Statement{{Resource: "**"}},
				},
			},
			want: true,
		},
		{
			name: "Single asterisk among other chars",
			policy: IAMRolePolicy{
				PolicyDocument: PolicyDocument{
					Statement: []Statement{{Resource: "A * s"}},
				},
			},
			want: true,
		},
		{
			name: "Single asterisk in Effect name",
			policy: IAMRolePolicy{
				PolicyDocument: PolicyDocument{
					Statement: []Statement{{Effect: "*"}},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := VerifyJSON(tt.policy); result != tt.want {
				t.Errorf("Expected: %t, got: %t", tt.want, result)
			}
		})
	}
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name   string
		policy IAMRolePolicy
		want   bool
	}{
		{
			name: "All necessary data",
			policy: IAMRolePolicy{
				PolicyName: "name",
				PolicyDocument: PolicyDocument{
					Version: "2012-10-17",
					Statement: []Statement{
						{
							Resource: "*",
							Effect:   "Allow",
							Action:   []string{"ok"},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Missing Version",
			policy: IAMRolePolicy{
				PolicyName: "name",
				PolicyDocument: PolicyDocument{
					Statement: []Statement{
						{
							Resource: "*",
							Effect:   "Allow",
							Action:   []string{"ok"},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "Missing statement",
			policy: IAMRolePolicy{
				PolicyName: "name",
				PolicyDocument: PolicyDocument{
					Version: "2012-10-17",
				},
			},
			want: false,
		},
		{
			name: "Missing Effect",
			policy: IAMRolePolicy{
				PolicyName: "name",
				PolicyDocument: PolicyDocument{
					Version: "2012-10-17",
					Statement: []Statement{
						{
							Resource: "*",
							Action:   []string{"ok"},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "Missing Action",
			policy: IAMRolePolicy{
				PolicyName: "name",
				PolicyDocument: PolicyDocument{
					Version: "2012-10-17",
					Statement: []Statement{
						{
							Resource: "*",
							Effect:   "Allow",
						},
					},
				},
			},
			want: false,
		},
		{
			name: "Missing resource",
			policy: IAMRolePolicy{
				PolicyName: "name",
				PolicyDocument: PolicyDocument{
					Version: "2012-10-17",
					Statement: []Statement{
						{
							Effect: "Allow",
							Action: []string{"ok"},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePolicy(tt.policy)
			result := (err == nil)
			if result != tt.want {
				t.Errorf("Expected: %t, got: %t", tt.want, result)
			}
		})
	}
}
