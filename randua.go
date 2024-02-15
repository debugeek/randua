package randua

import (
	"log"
	"math/rand"
)

type PLATFORM int

const (
	WINDOWS PLATFORM = iota
	LINUX   PLATFORM = iota
	MACOS   PLATFORM = iota
	IOS     PLATFORM = iota
	ANDROID PLATFORM = iota
)

type BROWSER int

const (
	CHROME      BROWSER = iota
	EDGE        BROWSER = iota
	SAFARI      BROWSER = iota
	FIREFOX     BROWSER = iota
	FIREFOX_IOS BROWSER = iota
)

func GetRandomUserAgent() string {
	platforms := []PLATFORM{WINDOWS, LINUX, MACOS, IOS, ANDROID}
	platform := platforms[rand.Intn(len(platforms))]

	var browsers []BROWSER
	switch platform {
	case WINDOWS:
		browsers = []BROWSER{CHROME, EDGE, FIREFOX}
	case LINUX:
		browsers = []BROWSER{CHROME, FIREFOX}
	case MACOS:
		browsers = []BROWSER{CHROME, EDGE, SAFARI, FIREFOX}
	case IOS:
		browsers = []BROWSER{CHROME, EDGE, SAFARI, FIREFOX}
	case ANDROID:
		browsers = []BROWSER{CHROME, EDGE, FIREFOX}
	}

	browser := browsers[rand.Intn(len(browsers))]

	safariVersion := BuildVersion(".", VersionRange{605, 632}, VersionRange{1, 30}, VersionRange{1, 15})
	chromeVersion := BuildVersion(".", VersionRange{105, 123}, "0", VersionRange{6099, 6226}, VersionRange{0, 283})
	edgeVersion := BuildVersion(".", VersionRange{105, 121}, "0", VersionRange{2088, 2277}, VersionRange{44, 167})
	firefoxVersion := BuildVersion(".", VersionRange{95, 122}, "0", VersionRange{0, 3})
	webkitVersion := safariVersion
	geckoVersion := firefoxVersion

	builder := NewUserAgentBuilder()

	{
		element := NewUserAgentElement("Mozilla/5.0")
		switch platform {
		case WINDOWS:
			element.AddComment("Windows NT 10.0")
		case LINUX:
			element.AddComment("Linux")
		case MACOS:
			element.AddComment("Macintosh")
			element.AddComment("Intel Mac OS X " + BuildVersion("_", VersionRange{12, 14}, VersionRange{0, 3}, VersionRange{0, 3}))
		case IOS:
			element.AddComment("iPhone")
			element.AddComment("CPU iPhone OS " + BuildVersion("_", VersionRange{11, 17}, VersionRange{0, 3}, VersionRange{0, 3}) + " like Mac OS X")
		case ANDROID:
			element.AddComment("Linux")
			element.AddComment("Android " + BuildVersion("_", VersionRange{12, 14}))
		}
		if browser == FIREFOX && platform != IOS {
			element.AddComment("rv:" + geckoVersion)
		}
		builder.AddElement(element)
	}

	{
		if platform != IOS && browser == FIREFOX {
			if platform == WINDOWS || platform == MACOS || platform == LINUX {
				// On Desktop, geckotrail is the fixed string "20100101"
				builder.AddElement(NewUserAgentElement("Gecko/20100101"))
			} else {
				// From Firefox 10 on mobile, geckotrail is the same as firefoxversion.
				builder.AddElement(NewUserAgentElement("Gecko/" + geckoVersion))
			}
		} else {
			element := NewUserAgentElement("AppleWebKit/" + webkitVersion)
			element.AddComment("KHTML, like Gecko")
			builder.AddElement(element)
		}
	}

	{
		switch browser {
		case CHROME:
			switch platform {
			case WINDOWS:
				builder.AddElement(NewUserAgentElement("Chrome/" + chromeVersion))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			case MACOS:
				builder.AddElement(NewUserAgentElement("Chrome/" + chromeVersion))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			case LINUX:
				builder.AddElement(NewUserAgentElement("Chrome/" + chromeVersion))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			case IOS:
				builder.AddElement(NewUserAgentElement("CriOS/" + chromeVersion))
				builder.AddElement(NewUserAgentElement("Mobile/15E148"))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			case ANDROID:
				builder.AddElement(NewUserAgentElement("Chrome/" + chromeVersion))
				builder.AddElement(NewUserAgentElement("Mobile"))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			default:
				log.Fatalf("Unsupported platform %v %v", platform, browser)
			}
		case EDGE:
			switch platform {
			case WINDOWS:
				builder.AddElement(NewUserAgentElement("Chrome/" + chromeVersion))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
				builder.AddElement(NewUserAgentElement("Edg/" + edgeVersion))
			case MACOS:
				builder.AddElement(NewUserAgentElement("Chrome/" + chromeVersion))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
				builder.AddElement(NewUserAgentElement("Edg/" + edgeVersion))
			case IOS:
				builder.AddElement(NewUserAgentElement("Version/17.0"))
				builder.AddElement(NewUserAgentElement("EdgiOS/" + edgeVersion))
				builder.AddElement(NewUserAgentElement("Mobile/15E148"))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			case ANDROID:
				builder.AddElement(NewUserAgentElement("Chrome/" + chromeVersion))
				builder.AddElement(NewUserAgentElement("Mobile"))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
				builder.AddElement(NewUserAgentElement("EdgA/" + edgeVersion))
			default:
				log.Fatalf("Unsupported platform %v %v", platform, browser)
			}
		case SAFARI:
			switch platform {
			case MACOS:
				builder.AddElement(NewUserAgentElement("Version/17.0"))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			case IOS:
				builder.AddElement(NewUserAgentElement("Version/17.0"))
				builder.AddElement(NewUserAgentElement("Mobile/15E148"))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			default:
				log.Fatalf("Unsupported platform %v %v", platform, browser)
			}
		case FIREFOX:
			if platform == IOS {
				builder.AddElement(NewUserAgentElement("FxiOS/" + firefoxVersion))
				builder.AddElement(NewUserAgentElement("Mobile/15E148"))
				builder.AddElement(NewUserAgentElement("Safari/" + safariVersion))
			} else {
				builder.AddElement(NewUserAgentElement("Firefox/" + firefoxVersion))
			}
		}
	}

	return builder.Build()
}
