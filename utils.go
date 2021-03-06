package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"

	"github.com/sergi/go-diff/diffmatchpatch"
)

var animals []string
var adjectives []string
var robots_txt string

type versionsInfo struct {
	VersionDate string
	VersionNum  int
}

func init() {
	rand.Seed(time.Now().Unix())
	animals = []string{"fawn", "peacock", "fox terrier", "civet", "musk deer", "seastar", "pigeon", "bull", "bumblebee", "crocodile", "flying squirrel", "elephant", "leopard seal", "baboon", "porcupine", "wolverine", "spider monkey", "vampire bat", "sparrow", "manatee", "possum", "swallow", "wildcat", "bandicoot", "labradoodle", "dragonfly", "tarsier", "snowy owl", "chameleon", "boykin", "puffin", "bison", "llama", "kitten", "stinkbug", "macaw", "parrot", "leopard cat", "prawn", "panther", "dogfish", "fennec", "frigatebird", "nurse shark", "turkey", "cockatoo", "neanderthal", "crow", "gopher", "reindeer", "earwig ", "anaconda", "panda", "ant", "silver fox", "collared peccary", "puppy", "common buzzard", "moose", "binturong", "wildebeest", "lovebird", "ferret", "persian", "marine toad", "woolly mammoth", "dalmatian", "bird", "umbrellabird", "kingfisher", "kangaroo", "stallion", "russian blue", "ostrich", "owl", "tawny owl", "affenpinscher", "caiman", "elephant seal", "octopus", "meerkat", "whale shark", "buck", "donkey", "red wolf", "mountain lion", "labrador retriever", "quetzal", "chamois", "sponge", "hamster", "orangutan", "sea urchin", "uakari", "doberman", "dormouse", "saint bernard", "bull shark", "ocelot", "sparrow", "spitz", "stoat", "snapping turtle", "dragonfly", "cougar", "alligator", "walrus", "glass lizard", "malayan tiger", "frog", "tiger", "armadillo", "chinchilla", "crab", "squid", "calf", "shrew", "dolphin", "royal penguin", "dingo", "turtle", "yellow-eyed penguin", "chimpanzee", "armadillo", "boa constrictor", "rabbit", "basking", "coyote", "chinook", "osprey", "sea lion", "fly", "sperm whale", "patas monkey", "tiffany", "mountain goat", "dodo", "worm", "cat", "warthog", "peccary", "shark", "pony", "monkey", "swan", "whippet", "beagle", "cougar", "anteater", "quail", "liger", "cheetah", "woodpecker", "egret", "eagle", "moose", "warthog", "honey bee", "snail", "stag beetle", "budgie", "molly", "magpie", "rhinoceros", "elephant", "kudu", "wombat", "tree frog", "goat", "lamb", "tropicbird", "human", "hog", "tang", "pool frog", "lemur", "ox", "dog", "lizard", "echidna", "great dane", "wallaby", "hawk", "dove", "jellyfish", "sloth", "macaque", "starfish", "sun bear", "guppy", "welsh corgi", "deer", "impala", "porpoise", "gazelle", "bichon", "seal", "wolf", "zebra shark", "mole", "narwhal", "hedgehog", "sheep", "horse", "bluetick", "colt", "spadefoot toad", "wildebeest", "piranha", "basenji", "mallard", "bull mastiff", "bear", "siberian husky", "bird", "badger", "red panda", "hammerhead", "rock hyrax", "kangaroo", "marsh frog", "mule", "weasel", "dogfish", "dachsbracke", "forest elephant", "oyster", "bat", "python", "coati", "platypus", "salamander", "cat", "caterpillar", "giraffe", "snake", "kid", "falcon", "robin", "guinea fowl", "tern", "sea lion", "dingo", "bolognese", "drake", "goose", "rat", "gentoo penguin", "iguana", "quail", "mouse", "horseshoe crab", "roebuck", "cattle dog", "fish", "poodle", "frog", "wolverine", "chinchilla", "bobcat", "grey seal", "hermit crab", "carolina", "shepherd", "gila monster", "snail", "mandrill", "leopard", "frilled lizard", "echidna", "rabbit", "bison", "barracuda", "foal", "ass", "eagle", "octopus", "avocet", "siamese", "dodo", "yorkie", "cockroach", "wallaroo", "tiger", "woodlouse", "glow worm", "fossa", "buffalo", "zorse", "albatross", "indri", "seahorse", "lemur", "louse", "ostrich", "humpback whale", "millipede", "fin whale", "joey", "pinscher", "dachshund", "proboscis monkey", "pelican", "chihuahua", "dogo", "indian rhinoceros", "wasp", "siberian", "raccoon dog", "yak", "stingray", "jack russel", "water vole", "foxhound", "sheep", "stork", "horse", "monkey", "woolly monkey", "waterbuck", "dunker", "cuscus", "ibis", "giraffe", "aardvark", "hummingbird", "grizzly bear", "otter", "pike", "minke whale", "pika", "stickbug", "pelican", "dugong", "bongo", "lemming", "shrimp", "piglet", "sabre-toothed tiger", "gemsbok", "tiger shark", "tuatara", "rottweiler", "elephant shrew", "ewe", "coati", "cichlid", "akita", "gharial", "thorny devil", "duck", "macaroni penguin", "steer", "setter", "pufferfish", "donkey", "mink", "macaw", "wolfhound", "white tiger", "ram", "ant", "rat", "marten", "galapagos tortoise", "crab", "horn shark", "blue whale", "koala", "starfish", "partridge", "sea squirt", "fire-bellied toad", "chipmunk", "ibex", "maltese", "clumber", "butterfly", "manta ray", "flamingo", "opossum", "parrot", "mastiff", "water buffalo", "okapi", "salmon", "tapir", "adelie", "killer whale", "lynx", "basilisk", "indian elephant", "oyster", "manta ray", "prairie dog", "chipmunk", "locust", "dog", "cottontop", "hyena", "spectacled bear", "oriole", "cobra", "pug", "monitor", "mandrill", "antelope", "chinstrap", "zebra", "chicken", "mule", "seal", "goat", "little penguin", "gull", "tasmanian devil", "caterpillar", "tamarin", "wrasse", "woodchuck", "otter", "penguin", "porcupine", "killer whale", "bear", "ferret", "dusky", "nightingale", "slow worm", "bat", "jaguar", "humboldt", "ermine", "saola", "emu", "lobster", "weasel", "nightingale", "hound", "bombay", "platypus", "electric eel", "asian elephant", "sea otter", "uguisu", "scorpion", "fox", "jerboa", "bengal tiger", "zebu", "lion", "zonkey", "ragdoll", "caracal", "bee", "kiwi", "puma", "common loon", "jackal", "malamute", "mayfly", "baboon", "terrier", "jellyfish", "vicuna", "penguin", "desert tortoise", "muskrat", "water dragon", "zebra", "malayan civet", "burmese", "orangutan", "himalayan", "pond skater", "howler monkey", "newt", "border collie", "cow", "bearded dragon", "fish", "barn owl", "puffin", "chin", "anteater", "beaver", "canary", "hamster", "sloth", "collie", "heron", "sea dragon", "gopher", "magpie", "king crab", "flounder", "opossum", "pademelon", "capybara", "boar", "leaf-tailed gecko", "turkey", "clown fish", "musk-ox", "bulldog", "pronghorn", "hercules beetle", "reindeer", "llama", "pygmy", "eskimo dog", "kinkajou", "komodo dragon", "cuttlefish", "cub", "bloodhound", "squirrel", "gander", "moorhen", "emu", "brown bear", "javanese", "birman", "harrier", "tortoise", "antelope", "gnu", "kingfisher", "wasp", "olm", "havanese", "canaan", "lizard", "indochinese tiger", "ocelot", "mist", "hare", "discus", "cony", "orca", "rooster", "ground hog", "silver dollar", "peacock", "akbash", "somali", "beaver", "maine coon", "mouse", "eland", "squirrel", "serval", "chimpanzee", "snowshoe", "toucan", "catfish", "lynx", "coyote", "bunny", "retriever", "fur seal", "cow", "balinese", "vulture", "coral", "leopard", "raccoon", "polar bear", "okapi", "kakapo", "whale", "sand lizard", "bonobo", "moray", "gila monster", "cormorant", "bracke", "camel", "markhor", "rockhopper", "neapolitan", "black bear", "roseate spoonbill", "woodpecker", "mountain lion", "crested penguin", "hippopotamus", "puma", "camel", "alligator", "guinea pig", "heron", "siberian tiger", "river dolphin", "axolotl", "argentino", "human", "mongoose", "drever", "quokka", "common frog", "elk", "wombat", "spider monkey", "civet", "sting ray", "panther", "gar", "lionfish", "snake", "crane", "newt", "raven", "tortoise", "fire ant", "chicadee", "common toad", "pig", "manatee", "centipede", "numbat", "river turtle", "falcon", "angelfish", "chamois", "rhinoceros", "shark", "flamingo", "pheasant", "ladybird", "grasshopper", "greyhound", "lemming", "pig", "marmoset", "eel", "yorkiepoo", "tiger salamander", "mosquito", "shih tzu", "quoll", "chick", "guanaco", "walrus", "badger", "ainu", "squid", "pekingese", "gerbil", "duck", "rattlesnake", "tapir", "lobster", "catfish", "mustang", "wallaby", "mongrel", "butterfly", "booby", "bush elephant", "fox", "rattlesnake", "cockroach", "tadpole", "lark", "ape", "pied tamarin", "mare", "tetra", "squirrel monkey", "elephant seal", "dhole", "cesky", "raccoon", "newfoundland", "marmoset", "stag", "bullfrog", "black bear", "crocodile", "lion", "barb", "wolf", "vervet monkey", "beetle", "polar bear", "pointer", "grizzly bear", "meerkat", "owl", "reptile", "fousek", "gibbon", "king penguin", "budgerigar", "swan", "hartebeest", "cassowary", "borneo elephant", "oryx", "alpaca", "gerbil", "chameleon", "galapagos penguin", "vulture", "barracuda", "insect", "fishing cat", "hen", "giant clam", "hare", "polecat", "fly", "chow chow", "yak", "seahorse", "spider", "eel", "burro", "brown bear", "boxer dog", "crane", "bandicoot", "hedgehog", "dromedary", "goose", "budgerigar", "dolphin", "kelpie dog", "highland cattle", "dormouse", "duckbill", "springbok", "mongoose", "bobcat", "water buffalo", "gecko", "hornet", "iguana", "wild boar", "koala", "guinea pig", "marmot", "skink", "deer", "filly", "barnacle", "tree toad", "leopard tortoise", "appenzeller", "doe", "gecko", "mole", "mynah bird", "mau", "gilla monster", "french bulldog", "termite", "salamander", "parakeet", "finch", "horned frog", "hippopotamus", "hummingbird", "cheetah", "albatross", "jaguar", "toad", "hyena", "gorilla", "skunk", "impala", "sea slug", "scorpion fish", "jackal", "skunk", "grouse", "sea turtle", "moth", "caribou", "dugong", "bighorn sheep", "ibizan hound", "gorilla", "puffer fish", "chicken", "komodo dragon", "buffalo"}
	adjectives = []string{"lulling", "vile", "foul", "cheerful", "messy", "dreadful", "uneven", "stinky", "young", "sparkling", "sweltering", "verdant", "hideous", "friendly", "blistering", "rambunctious", "carefree", "fat", "sloppy", "gloomy", "awful", "anemic", "minute", "stiff", "benevolent", "ceaseless", "large", "quick", "round", "glassy", "rusty", "scarce", "odd", "shining", "even", "dowdy", "solemn", "scorching", "brief", "rotten", "new", "plush", "cozy", "meandering", "apologetic", "nimble", "busy", "strong", "great", "brilliant", "piercing", "creepy", "miniature", "narrow", "whimsical", "fantastic", "cowardly", "disgusting", "marvelous", "snug", "stern", "stingy", "angry", "spiky", "cheeky", "gorgeous", "mysterious", "flat", "clever", "charming", "dismal", "meek", "somber", "sour", "thin", "beautiful", "stubborn", "crazy", "challenging", "gaunt", "salty", "indifferent", "huge", "daring", "awkward", "picturesque", "copious", "glowing", "truthful", "rude", "petite", "cranky", "ornery", "brazen", "modest", "purring", "filthy", "rotund", "short", "splendid", "hasty", "deafening", "crawling", "furious", "tall", "mute", "ghastly", "still", "arrogant", "crabby", "haughty", "curly", "voiceless", "hot", "courageous", "late", "microscopic", "vast", "stifling", "good", "disrespectful", "bashful", "moaning", "towering", "adventurous", "idyllic", "careless", "shocking", "erratic", "heavy", "square", "hard", "jagged", "gullible", "hushed", "revolting", "content", "thrifty", "sluggish", "eternal", "greedy", "pleasant", "small", "demanding", "greasy", "enormous", "hostile", "terrible", "chilly", "speedy", "massive", "loud", "puny", "striking", "clumsy", "soaring", "brave", "fast", "delicious", "effortless", "bland", "thick", "little", "ancient", "silent", "wonderful", "stuffed", "grimy", "bitter", "muggy", "shallow", "ridiculous", "absent-minded", "fuzzy", "peculiar", "mangy", "wide", "kind", "squeaky", "screeching", "silly", "squealing", "spoiled", "gigantic", "happy", "steep", "ingenious", "modern", "juicy", "gentle", "medium", "brawny", "curved", "lumpy", "afraid", "amusing", "thundering", "cooing", "oppressive", "swollen", "grave", "sturdy", "average", "proud", "rancid", "absurd", "entertaining", "annoyed", "fussy", "precise", "subtle", "gilded", "slow", "delinquent", "nervous", "hopeful", "rich", "adequate", "shrill", "plump", "freezing", "nasty", "endless", "lavish", "worried", "courteous", "bulky", "fair", "diminutive", "groggy", "miserable", "horrid", "crooked", "monstrous", "superb", "contrary", "lazy", "fidgety", "menacing", "swift", "stale", "quarrelsome", "quiet", "askew", "tough", "simple", "sweet", "hardworking", "frosty", "whispering", "famished", "crispy", "caring", "capable", "tiny", "immense", "startled", "lovely", "high-pitched", "tasteless", "decrepit", "tense", "lousy", "straight", "excited", "ugly", "stunning", "parched", "wild", "ripe", "lonely", "optimistic", "obnoxious", "cavernous", "different", "harsh", "creaky", "grand", "difficult", "temporary", "eccentric", "muffled", "alert", "delicate", "timid", "infamous", "enchanting", "anxious", "humble", "edgy", "severe", "repulsive", "desolate", "sleepy", "slimy", "irritable", "vigilant", "generous", "rapid", "old-fashioned", "hilly", "easy", "righteous", "joyful", "surprised", "starving", "big", "early", "compassionate", "moody", "perpetual", "dishonest", "serious", "foolish", "soft", "old", "scared", "mighty", "trendy", "curious", "hissing", "savage", "dense", "steaming", "broad", "slick", "creative", "icy", "adorable", "slight", "terrified", "intense", "noisy", "cautious", "sizzling", "blithe", "fluttering", "faint", "delighted", "smelly", "lively", "frightened", "gauzy", "long", "strict", "bored", "calm", "melodic", "spicy", "relaxed", "triangular", "dull", "wise", "dangerous", "smooth", "cruel", "creeping", "dawdling", "intimidating", "exhausted", "deep", "tasty", "obtuse", "graceful", "tranquil", "raspy", "selfish", "sullen", "malicious", "ecstatic", "wrinkly", "opulent", "polite", "fetid", "husky", "prudent", "skinny", "tricky", "impatient", "loyal", "fresh"}
	about_page = `# Cowyo...

_...is the Collection of Online Words You Open._

This tool is supposed to make sharing online notes and lists fast and easy. To jot a note, simply load the page at [` + "`" + `/` + "`" + `](/) and write. The url will redirect to an easy-to-remember name that you can use to reload the page at anytime, anywhere. (You can use any url you want too: [` + "`" + `/AnythingYouWant` + "`" + `](/AnythingYouWant)). No need to press save, it will automatically save when you stop writing.

You can also write your notes in [Markdown](https://daringfireball.net/projects/markdown/) and then render your page by adding ` + "`" + `/view` + "`" + `. For example, the page ` + "`" + `/about` + "`" + ` is rendered at [` + "`" + `/about/view` + "`" + `](/about/view).

If you are writing a list and you want to tick off things really easily, just add ` + "`" + `/list` + "`" + `. For example, after editing [` + "`" + `/grocery` + "`" + `](/grocery), goto [` + "`" + `/grocery/list` + "`" + `](/grocery/list). In this page, whatever you click on will be striked through and moved to the end. This is helpful if you write a grocery list and then want to easily delete things from it.

Base64 images are supported [in img tags](https://stackoverflow.com/questions/1207190/embedding-base64-images). Math is supported using [Katex](https://github.com/Khan/KaTeX). For example, ` + "`" + `&#36;\frac{1}{2}&#36;` + "`" + ` becomes $\frac{1}{2}$ and ` + "`" + `&#36;&#36;E=mc^2&#36;&#36;` + "`" + ` becomes $$E=mc^2$$

Be cautious about writing sensitive information in the notes as anyone with the URL has access to it. For more information, or if you'd like to edit the code, [use the Github](https://github.com/schollz/cowyo). If you'd like help or find a bug, please submit [an issue](https://github.com/schollz/cowyo/issues) or <a href="https://twitter.com/intent/tweet?screen_name=zack_118" class="twitter-mention-button" data-related="zack_118">Tweet me @zack_118</a>.

Have fun.

**Powered by Raspberry Pi, Go, and NGINX**

![Raspberry Pi](/static/img/raspberrypi.png) ![Go Mascot](/static/img/gomascot.png) ![Nginx](/static/img/nginx.png)`

	robots_txt = `User-agent: *
Disallow: /`

}

func randomAnimal() string {
	return strings.Replace(strings.Title(animals[rand.Intn(len(animals)-1)]), " ", "", -1)
}

func randomAdjective() string {
	return strings.Replace(strings.Title(adjectives[rand.Intn(len(adjectives)-1)]), " ", "", -1)
}

func randomAlliterateCombo() (combo string) {
	combo = ""
	for {
		animal := randomAnimal()
		adjective := randomAdjective()
		if animal[0] == adjective[0] {
			combo = adjective + animal
			break
		}
	}
	return
}

func contentType(filename string) string {
	switch {
	case strings.Contains(filename, ".css"):
		return "text/css"
	case strings.Contains(filename, ".jpg"):
		return "image/jpeg"
	case strings.Contains(filename, ".png"):
		return "image/png"
	case strings.Contains(filename, ".js"):
		return "application/javascript"
	}
	return "text/html"
}

var about_page string

func diffRebuildtexts(diffs []diffmatchpatch.Diff) []string {
	text := []string{"", ""}
	for _, myDiff := range diffs {
		if myDiff.Type != diffmatchpatch.DiffInsert {
			text[0] += myDiff.Text
		}
		if myDiff.Type != diffmatchpatch.DiffDelete {
			text[1] += myDiff.Text
		}
	}
	return text
}

func getImportantVersions(p CowyoData) []versionsInfo {
	m := map[int]int{}
	dmp := diffmatchpatch.New()
	lastText := ""
	lastTime := time.Now().AddDate(0, -1, 0)
	for i, diff := range p.Diffs {
		seq1, _ := dmp.DiffFromDelta(lastText, diff)
		texts_linemode := diffRebuildtexts(seq1)
		rebuilt := texts_linemode[len(texts_linemode)-1]
		parsedTime, _ := time.Parse(time.ANSIC, p.Timestamps[i])
		duration := parsedTime.Sub(lastTime)
		m[i] = int(duration.Seconds())
		if i > 0 {
			m[i-1] = m[i]
		}
		// On to the next one
		lastText = rebuilt
		lastTime = parsedTime
	}

	// Sort in order of decreasing diff times
	n := map[int][]int{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	// Get the top 4 biggest diff times
	var importantVersions []int
	var r []versionsInfo
	for _, k := range a {
		for _, s := range n[k] {
			if s != 0 && s != len(n) {
				fmt.Printf("%d, %d\n", s, k)
				importantVersions = append(importantVersions, s)
				if len(importantVersions) > 10 {
					sort.Ints(importantVersions)
					for _, nn := range importantVersions {
						r = append(r, versionsInfo{p.Timestamps[nn], nn})
					}
					return r
				}
			}
		}
	}
	sort.Ints(importantVersions)
	for _, nn := range importantVersions {
		r = append(r, versionsInfo{p.Timestamps[nn], nn})
	}
	return r
}

func rebuildTextsToDiffN(p CowyoData, n int) string {
	dmp := diffmatchpatch.New()
	lastText := ""
	for i, diff := range p.Diffs {
		seq1, _ := dmp.DiffFromDelta(lastText, diff)
		texts_linemode := diffRebuildtexts(seq1)
		rebuilt := texts_linemode[len(texts_linemode)-1]
		if i == n {
			return rebuilt
		}
		lastText = rebuilt
	}
	return "ERROR"
}
