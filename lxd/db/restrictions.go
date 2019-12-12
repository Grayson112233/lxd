package db

func (c *ClusterTx) RestrictionsForProject(project_name string) (string, error) {
	stmt := "SELECT restrictions_list FROM restrictions WHERE project_name=?"

	restrictions, err := query.SelectObjects(stmt, project_name)
	if err != nil {
		return nil, err
	}

	restrictions := map[string]bool{}
	for _, restriction := range restrictions {
		restrictions[restriction.value] = restriction.enabled
	}
	return restrictions, nil
}