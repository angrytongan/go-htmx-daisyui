package nav

import "slices"

type NavLink struct {
	Label  string
	Href   string
	Active bool
}

var navigationLinks = []NavLink{
	{Label: "root", Href: "/", Active: false},
	{Label: "icons", Href: "/icons", Active: false},
	{Label: "badges", Href: "/badges", Active: false},
	{Label: "fragments", Href: "/template-fragments", Active: false},
}

func MakeLinks(activeHref string) []NavLink {
	nav := slices.Clone(navigationLinks)
	for i, v := range nav {
		nav[i].Active = v.Href == activeHref
	}

	return nav
}
