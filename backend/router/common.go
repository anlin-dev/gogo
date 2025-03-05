package router

func commonGroups() []CommonRouter {

	return []CommonRouter{
		&DashboardRouter{},
	}
}
