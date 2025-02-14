package casbin

func (casbinConfig *CasbinConfig) AddPolicy(params ...string) error {
	casbinConfig.Enforce.AddPolicy(params)

	return casbinConfig.Enforce.SavePolicy()
}

func (casbinConfig *CasbinConfig) AddGroupingPolicy(params ...string) error {
	casbinConfig.Enforce.AddGroupingPolicy(params)

	return casbinConfig.Enforce.SavePolicy()
}

func (casbinConfig *CasbinConfig) UpdatePolicy(current []string, new []string) error {
	casbinConfig.Enforce.UpdatePolicy(current, new)

	return casbinConfig.Enforce.SavePolicy()
}

func GetPermissionsForUser(value string) ([][]string, error) {
	permissions, err := Casbin.Enforce.GetImplicitPermissionsForUser(value)
	if err != nil {
		return nil, err
	}

	return permissions, nil
}
