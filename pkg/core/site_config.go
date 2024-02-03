package core

type SiteConfig struct {
	Name        string
	Description string
	MainNav     []NavLink
	Links       []SiteLink
	Projects    []ProjectPreview
}

type SiteLink struct {
	Name string
	Url  string
}

func Site() SiteConfig {
	return SiteConfig{
		Name:        "jon",
		Description: "My personal website",
		MainNav:     navlinks(),
		Links:       links(),
		Projects:    projects(),
	}
}

func links() []SiteLink {
	return []SiteLink{
		{
			Name: "GitHub",
			Url:  "https://github.com/jontitorr",
		},
		{
			Name: "Twitter",
			Url:  "https://twitter.com/jontitorr",
		},
	}
}

type NavLink struct {
	Name string
	Url  string
}

func navlinks() []NavLink {
	return []NavLink{
		{
			Name: "Home",
			Url:  "/",
		},
		{
			Name: "Projects",
			Url:  "/projects",
		},
	}
}

type ProjectPreview struct {
	Name string
	Html string
	Url  string
}

func projects() []ProjectPreview {
	return []ProjectPreview{
		{
			Name: "Ekizu",
			Html: `A modular <a href="https://discord.com">Discord</a> API wrapper for C++ applications.`,
			Url:  "https://github.com/jontitorr/ekizu",
		},
		{
			Name: "Shiki",
			Html: `A chat application with a backend/frontend similar to Discord's. Built with Rust and Svelte.`,
			Url:  "https://shiki.space",
		},
		{
			Name: "YTunes",
			Html: `An Alexa Skill that plays music from YouTube. Built with TypeScript and deployed using AWS Lambda.`,
			Url:  "https://github.com/jontitorr/ytunes",
		},
		{
			Name: "Saber",
			Html: `An <a href="https://discord.com">Discord</a> bot that can be used for moderation. Built with C++.`,
			Url:  "https://github.com/jontitorr/saber",
		},
	}
}
