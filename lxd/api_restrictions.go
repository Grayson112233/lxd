package main

var restrictionsCmd = APIEndpoint {
	Path: "restrictions",

	Get: APIEndpointAction{Handler: restrictionsGet, AccessHandler: AllowAuthenticated},
}

var restrictionsProjectCmd = APIEndpoint {
	Path: "restrictions/{project_name}",

	Get: APIEndpointAction{Handler: restrictionsProjectGet, AccessHandler: AllowAuthenticated},
	Put: APIEndpointAction{Handler: restrictionsProjectPut, AccessHandler: AllowAuthenticated},
}

var restrictionsKeyCmd = APIEndpoint {
	Path: "restrictions/{project_name}/{restriction_key}",

	Get: APIEndpointAction{Handler: restrictionKeyGet, AccessHandler: AllowAuthenticated},
	Put: APIEndpointAction{Handler: restrictionKeyPut, AccessHandler: AllowAuthenticated},
}

/**
Return a list of restrictions features available for LXD
*/
func restrictionsGet(d *Daemon, r *http.Request) response.Response {
	restrictions := []api.RestrictionsList{}
	return restrictions
}

/**
Return all current restrictions for a given project by JSON format
*/
func restrictionsProjectGet(d *Daemon, r *http.Request) response.Response {
	// Get project name
	project_name := projectParam(r)

	var result map[string][]string 

	err := d.cluster.Transaction(func(tx *db.ClusterTx) error {
		restrictionsEnabled, err := projecttx.RestrictionsEnabled(project_name)

		if err != nil {
			return errors.Wrap(err, "Check project restrictions enabled")
		}

		if !restrictionsEnabled {
			return response.Response("Restrictions have not been enabled for this project!")
		}

	})
}

/*
Set/Update all restrictiions for a given project. 
Pass in JSON object defining all restrictions and values to set
*/
func restrictionsProjectPut(d *Daemon, r *http.Request) response.Response {

}

/*
Return the value of specified restriction key for a project
*/
func restrictionKeyGet(d *Daemon, r *http.Request) response.Response {

}

/*
Set the value of a specified restriction key for a project
*/
func restrictionKeyPut(d *Daemon, r *http.Request) response.Response {

}