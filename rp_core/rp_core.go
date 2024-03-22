package rp_core

type RolePolicy struct {
	PolicyName     string
	PolicyDocument struct {
		Version   *string
		Statement []struct {
			Sid      string
			Effect   string
			Action   []string
			Resource string
		}
	}
}
