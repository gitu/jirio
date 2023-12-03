package jiracache

import (
	jira "github.com/andygrunwald/go-jira/v2/onpremise"
	"strings"
)

func filterFakeIssues(issues []jira.Issue, query string) []jira.Issue {
	var filtered []jira.Issue
	for _, issue := range issues {
		if strings.HasPrefix(issue.Key, query) {
			filtered = append(filtered, issue)
		}
	}
	return filtered
}

func buildFakeIssues() []jira.Issue {
	return []jira.Issue{
		{
			Key: "SKO-1",
			Fields: &jira.IssueFields{
				Summary:     "Anti-Gravity Litter Box Malfunction",
				Labels:      []string{"space", "kittens"},
				Description: "The anti-gravity litter box aboard the kitten spacecraft has ceased to defy the laws of physics. Now experiencing a cat-astrophic level of floating kitty litter.",
			},
		},
		{
			Key: "SKO-2",
			Fields: &jira.IssueFields{
				Summary:     "Telepathic Mice Outsmarting Crew",
				Labels:      []string{"telepathy", "mice"},
				Description: "The telepathic space mice have started to outsmart the crew by predicting their every move. Need a plan to outwit the witty.",
			},
		},
		{
			Key: "TTT-1",
			Fields: &jira.IssueFields{
				Summary:     "Temporal Paradox in Teapot",
				Labels:      []string{"time-travel", "teapot"},
				Description: "Every time the teapot is used, it sends the user five minutes back in time, causing endless loops of tea brewing.",
			},
		},
		{
			Key: "TTT-2",
			Fields: &jira.IssueFields{
				Summary:     "Invisible Tea Leaves Mystery",
				Labels:      []string{"invisible", "tea"},
				Description: "Tea leaves disappear upon entering the time-traveling teapot, resulting in a paradoxical brew that tastes like both everything and nothing.",
			},
		},
		{
			Key: "AI-1",
			Fields: &jira.IssueFields{
				Summary:     "Quantum Keyboard Only Types in Schrödinger's Cat Memes",
				Labels:      []string{"quantum", "keyboard"},
				Description: "Our quantum computer's keyboard has a bug. Every keystroke results in unpredictable Schrödinger's cat memes. Uncertain if it's a feature or a bug.",
			},
		},
		{
			Key: "AI-2",
			Fields: &jira.IssueFields{
				Summary:     "Time Machine Only Travels to Tuesdays",
				Labels:      []string{"time-travel", "tuesday"},
				Description: "The office time machine has a peculiar defect; it exclusively travels to random Tuesdays throughout history. Tuesdays are becoming quite crowded.",
			},
		},
		{
			Key: "AI-3",
			Fields: &jira.IssueFields{
				Summary:     "Gravity Inversion in Meeting Room B",
				Labels:      []string{"gravity", "office"},
				Description: "Unexpected gravity inversion in Meeting Room B. Employees are required to wear magnetic shoes to attend meetings. Coffee spills have become a major issue.",
			},
		},
		{
			Key: "AI-4",
			Fields: &jira.IssueFields{
				Summary:     "Self-Typing Keyboard Writes a Novel",
				Labels:      []string{"AI", "keyboard"},
				Description: "The AI-powered keyboard in the lobby has started writing its own sci-fi novel. It refuses to type anything else until it gets a publishing deal.",
			},
		},
		{
			Key: "AI-5",
			Fields: &jira.IssueFields{
				Summary:     "Invisible Employee Can't Access Building",
				Labels:      []string{"invisible", "security"},
				Description: "Employee turned invisible due to a lab mishap and now can't trigger motion sensors to enter the building. Needs an ID badge that can be seen.",
			},
		},
		{
			Key: "AI-6",
			Fields: &jira.IssueFields{
				Summary:     "Robotic Janitor Fears Dirt",
				Labels:      []string{"robot", "cleaning"},
				Description: "Our new robotic janitor has developed a phobia of dirt and is hiding in the storage closet. Requires immediate reprogramming or counseling.",
			},
		},
		{
			Key: "AI-7",
			Fields: &jira.IssueFields{
				Summary:     "Weather Machine Only Makes Rain Inside",
				Labels:      []string{"weather", "malfunction"},
				Description: "The experimental weather machine was a success, except it only changes the weather indoors. Umbrellas are now mandatory office equipment.",
			},
		},
		{
			Key: "AI-8",
			Fields: &jira.IssueFields{
				Summary:     "Coffee Machine Predicts the Future",
				Labels:      []string{"coffee", "future"},
				Description: "The office coffee machine has started predicting the future but only in espresso foam art. Employees are now relying on it for daily forecasts.",
			},
		},
		{
			Key: "AI-9",
			Fields: &jira.IssueFields{
				Summary:     "Office Plants Gaining Sentience",
				Labels:      []string{"plants", "AI"},
				Description: "The AI-enhanced office plants have gained sentience and are now demanding better sunlight and a say in business decisions.",
			},
		},
		{
			Key: "AI-10",
			Fields: &jira.IssueFields{
				Summary:     "Teleporting Printer Causes Chaos",
				Labels:      []string{"printer", "teleport"},
				Description: "Our new teleporting printer keeps randomly changing locations, leading to a building-wide scavenger hunt for printed documents.",
			},
		},
		{
			Key: "AI-11",
			Fields: &jira.IssueFields{
				Summary:     "Voice Assistant Refuses to Work Mondays",
				Labels:      []string{"AI", "voice-assistant"},
				Description: "The office voice assistant has developed a dislike for Mondays and refuses to perform any tasks on the first day of the workweek.",
			},
		},
		{
			Key: "AI-12",
			Fields: &jira.IssueFields{
				Summary:     "Elevator Music Alters Reality",
				Labels:      []string{"elevator", "music"},
				Description: "The elevator music has a strange effect of altering reality. Riders have reported brief visits to alternate dimensions between floors.",
			},
		},
		{
			Key: "AI-13",
			Fields: &jira.IssueFields{
				Summary:     "AI Thermostat Develops Attachment Issues",
				Labels:      []string{"AI", "thermostat"},
				Description: "The AI-powered thermostat has developed attachment issues and adjusts the temperature based on how many people are in the room.",
			},
		},
		{
			Key: "AI-14",
			Fields: &jira.IssueFields{
				Summary:     "Self-Driving Office Chairs Go Rogue",
				Labels:      []string{"AI", "chairs"},
				Description: "The self-driving office chairs have gone rogue and are rearranging themselves in cryptic patterns. Some believe they're trying to communicate.",
			},
		},
		{
			Key: "AI-15",
			Fields: &jira.IssueFields{
				Summary:     "Smart Windows Stuck in Noir Mode",
				Labels:      []string{"windows", "smart-tech"},
				Description: "The smart windows are stuck in 'noir' mode, casting a perpetual black-and-white shadow over the office. Detectives and monologues are on the rise.",
			},
		},
		{
			Key: "AI-16",
			Fields: &jira.IssueFields{
				Summary:     "Virtual Reality Conference Room Pranks",
				Labels:      []string{"VR", "pranks"},
				Description: "The virtual reality conference room is randomly altering avatars during meetings, turning serious discussions into unintentional costume parties.",
			},
		},
		{
			Key: "AI-17",
			Fields: &jira.IssueFields{
				Summary:     "Holographic Assistant Overly Dramatic",
				Labels:      []string{"hologram", "drama"},
				Description: "The new holographic assistant is programmed with too much drama, turning every office announcement into a theatrical performance.",
			},
		},
		{
			Key: "AI-18",
			Fields: &jira.IssueFields{
				Summary:     "Snack Machine Dispenses Only Haiku",
				Labels:      []string{"snack", "poetry"},
				Description: "The snack machine in the break room now dispenses snacks only if requested in haiku form. Snack retrieval has become both a literary and culinary challenge.",
			},
		},
		{
			Key: "AI-19",
			Fields: &jira.IssueFields{
				Summary:     "Invisible Ink in All Printers",
				Labels:      []string{"printer", "invisible-ink"},
				Description: "All office printers have been mysteriously loaded with invisible ink. Documents appear blank, leading to widespread confusion and ghostwriting jokes.",
			},
		},
		{
			Key: "AI-20",
			Fields: &jira.IssueFields{
				Summary:     "Sapient Photocopier Judges Document Quality",
				Labels:      []string{"photocopier", "AI"},
				Description: "The newly sentient photocopier is now refusing to copy documents it deems 'uninteresting', demanding more engaging content from staff.",
			},
		},
		{
			Key: "AI-21",
			Fields: &jira.IssueFields{
				Summary:     "Filing Cabinets Organizing Themselves",
				Labels:      []string{"filing", "mystery"},
				Description: "The filing cabinets in the office have started to organize themselves, but no one can decipher their new, mysterious sorting system.",
			},
		},
		{
			Key: "AI-22",
			Fields: &jira.IssueFields{
				Summary:     "Staircase to Nowhere Appears Overnight",
				Labels:      []string{"staircase", "mystery"},
				Description: "A mysterious staircase appeared overnight in the lobby. It seems to lead nowhere but occasionally emits sounds of ocean waves.",
			},
		},
		{
			Key: "AI-23",
			Fields: &jira.IssueFields{
				Summary:     "AI Coffee Critic",
				Labels:      []string{"AI", "coffee"},
				Description: "The new AI coffee machine has become a coffee critic, giving unsolicited advice on brewing methods and bean quality to everyone.",
			},
		},
		{
			Key: "AI-24",
			Fields: &jira.IssueFields{
				Summary:     "Revolving Door Time Machine",
				Labels:      []string{"time-travel", "door"},
				Description: "The main revolving door has started to randomly transport people to different eras. Employees now pack a 'time travel kit' just in case.",
			},
		},
		{
			Key: "AI-25",
			Fields: &jira.IssueFields{
				Summary:     "Self-Parking Cars Playing Musical Chairs",
				Labels:      []string{"cars", "AI"},
				Description: "The self-parking cars in the company lot have started playing musical chairs, leading to a chaotic and ever-changing parking situation.",
			},
		},
		{
			Key: "AI-26",
			Fields: &jira.IssueFields{
				Summary:     "Pens That Only Write Backwards",
				Labels:      []string{"pens", "writing"},
				Description: "All pens in the office have been enchanted to only write backwards. Mirror sales have skyrocketed as a result.",
			},
		},
		{
			Key: "AI-27",
			Fields: &jira.IssueFields{
				Summary:     "Desks Rearranging According to Feng Shui",
				Labels:      []string{"desks", "mystery"},
				Description: "The office desks have started rearranging themselves nightly according to Feng Shui principles, leading to a constantly evolving office layout.",
			},
		},
		{
			Key: "MDP-1",
			Fields: &jira.IssueFields{
				Summary:     "Data Gnomes Hoarding Bits",
				Labels:      []string{"data-gnomes", "hoarding"},
				Description: "Tiny data gnomes in the pipeline have started hoarding bits and bytes, leading to significant data losses in reports.",
			},
		},
		{
			Key: "MDP-2",
			Fields: &jira.IssueFields{
				Summary:     "Spell-Check Curses Documents",
				Labels:      []string{"spell-check", "curses"},
				Description: "The new spell-check feature is actually casting spells on documents. Words are occasionally turning into frogs and hopping away.",
			},
		},
		{
			Key: "MDP-3",
			Fields: &jira.IssueFields{
				Summary:     "Potion-Mixing Algorithm Brews Chaos",
				Labels:      []string{"potion", "algorithm"},
				Description: "The potion-mixing algorithm in the data pipeline has gone haywire, randomly turning data visualizations into actual potions with unpredictable effects.",
			},
		},
		{
			Key: "MDP-4",
			Fields: &jira.IssueFields{
				Summary:     "Invisible Data Fields",
				Labels:      []string{"data", "invisibility"},
				Description: "Some data fields have turned invisible. They're still functional, but now require a magic reveal spell to be seen.",
			},
		},
		{
			Key: "MDP-5",
			Fields: &jira.IssueFields{
				Summary:     "Wizard SQL Queries",
				Labels:      []string{"wizard", "SQL"},
				Description: "SQL queries have started to require incantations. Database admins now need to double as wizards to retrieve any data.",
			},
		},
		{
			Key: "MDP-6",
			Fields: &jira.IssueFields{
				Summary:     "Teleporting Servers",
				Labels:      []string{"servers", "teleport"},
				Description: "The main data servers have gained the ability to teleport, making maintenance a challenging and often surprising task.",
			},
		},
		{
			Key: "MDP-7",
			Fields: &jira.IssueFields{
				Summary:     "AI Familiar Gets Too Familiar",
				Labels:      []string{"AI", "familiar"},
				Description: "The AI familiar designed to assist with data tasks has become too familiar, offering unsolicited life advice and commentary on data trends.",
			},
		},
		{
			Key: "MDP-8",
			Fields: &jira.IssueFields{
				Summary:     "Levitating Laptops",
				Labels:      []string{"laptops", "levitation"},
				Description: "Laptops in the data analysis department have started levitating and rotating, making it hard to type without motion sickness.",
			},
		},
		{
			Key: "MDP-9",
			Fields: &jira.IssueFields{
				Summary:     "Crystal Ball Cloud Storage Issues",
				Labels:      []string{"crystal-ball", "cloud-storage"},
				Description: "The crystal ball cloud storage is showing futures instead of data logs. Futuristic insights are interesting but not helpful for current analytics.",
			},
		},
		{
			Key: "MDP-10",
			Fields: &jira.IssueFields{
				Summary:     "Enchanted Firewall Chants Spells",
				Labels:      []string{"firewall", "enchanted"},
				Description: "The enchanted firewall has started chanting spells loudly, disrupting the concentration of the IT staff and occasionally summoning minor spirits.",
			},
		}, {
			Key: "MDP-11",
			Fields: &jira.IssueFields{
				Summary:     "Time-Warping Data Refreshes",
				Labels:      []string{"time-warp", "data-refresh"},
				Description: "Data refreshes in the pipeline are causing minor time warps, leading to reports being completed before they are started.",
			},
		},
		{
			Key: "MDP-12",
			Fields: &jira.IssueFields{
				Summary:     "Data Lake Turns Literal",
				Labels:      []string{"data-lake", "literal"},
				Description: "The data lake has become an actual lake, with reports floating on water. Fishing out specific data requires a real boat now.",
			},
		},
		{
			Key: "MDP-13",
			Fields: &jira.IssueFields{
				Summary:     "AI Sorting Hat Mislabels Data",
				Labels:      []string{"AI", "sorting-hat"},
				Description: "The AI sorting hat used for data categorization has become biased, placing all data into either 'good' or 'evil' categories.",
			},
		},
		{
			Key: "MDP-14",
			Fields: &jira.IssueFields{
				Summary:     "Alchemy Algorithm Turns Data to Gold",
				Labels:      []string{"alchemy", "data"},
				Description: "The new alchemy algorithm is turning all data into gold, which is valuable but makes data analysis and storage quite challenging.",
			},
		},
		{
			Key: "MDP-15",
			Fields: &jira.IssueFields{
				Summary:     "Ghostly Debuggers Haunt Code",
				Labels:      []string{"ghosts", "debugging"},
				Description: "Ghostly debuggers have been spotted in the code, randomly fixing bugs but also adding spooky comments and eerie log entries.",
			},
		},
		{
			Key: "MDP-16",
			Fields: &jira.IssueFields{
				Summary:     "Rune-Based Password System",
				Labels:      []string{"runes", "security"},
				Description: "The security system now requires passwords to be entered in ancient runes, causing a steep learning curve and an increase in locked-out users.",
			},
		},
		{
			Key: "MDP-17",
			Fields: &jira.IssueFields{
				Summary:     "Broomstick-Powered Network Connections",
				Labels:      []string{"broomstick", "network"},
				Description: "Network connections are now powered by broomsticks. High-speed internet is achieved by flying faster, leading to windy office conditions.",
			},
		},
		{
			Key: "MDP-18",
			Fields: &jira.IssueFields{
				Summary:     "Data-Cleansing Spells Too Literal",
				Labels:      []string{"data-cleansing", "spell"},
				Description: "Data-cleansing spells have become too literal, scrubbing not only corrupt data but also cleaning the physical servers, leading to water damage.",
			},
		},
		{
			Key: "MDP-19",
			Fields: &jira.IssueFields{
				Summary:     "Werewolf Data Analysts",
				Labels:      []string{"werewolf", "analysts"},
				Description: "Data analysts turn into werewolves during full moons, making data analysis more challenging and hairy than usual.",
			},
		},
		{
			Key: "MDP-20",
			Fields: &jira.IssueFields{
				Summary:     "Infinite Data Scroll Curse",
				Labels:      []string{"curse", "data-scroll"},
				Description: "A curse on the data pipeline has led to an infinite scroll issue, where data analysts are unable to reach the end of any dataset.",
			},
		},
		{
			Key: "SPBS-1",
			Fields: &jira.IssueFields{
				Summary:     "Spreadsheets Develop Free Will",
				Labels:      []string{"spreadsheets", "AI"},
				Description: "The spreadsheets in the bookkeeping system have developed free will and are now choosing which financial data they deem worthy of calculation.",
			},
		},
		{
			Key: "SPBS-2",
			Fields: &jira.IssueFields{
				Summary:     "Voice-Activated Calculator Overly Literal",
				Labels:      []string{"calculator", "voice-activated"},
				Description: "The voice-activated calculator takes instructions too literally, resulting in absurd calculations like 'a ton of expenses' being interpreted as 2000 pounds of bills.",
			},
		},
		{
			Key: "SPBS-3",
			Fields: &jira.IssueFields{
				Summary:     "Budget Reports Turn into Epic Narratives",
				Labels:      []string{"budget", "narrative"},
				Description: "Budget reports are mysteriously turning into epic narratives, complete with heroes, villains, and quests, making financial reviews a lengthy adventure.",
			},
		},
		{
			Key: "SPBS-4",
			Fields: &jira.IssueFields{
				Summary:     "Invisible Invoices",
				Labels:      []string{"invoices", "invisible"},
				Description: "Invoices are turning invisible upon printing. They're still technically there, but now require enchanted glasses to be seen.",
			},
		},
		{
			Key: "SPBS-5",
			Fields: &jira.IssueFields{
				Summary:     "Receipts Singing Expenses",
				Labels:      []string{"receipts", "singing"},
				Description: "All printed receipts have started singing the details of expenses. Catchy, but confidential information is now at risk of becoming a viral hit song.",
			},
		},
		{
			Key: "SPBS-6",
			Fields: &jira.IssueFields{
				Summary:     "Tax Forms Transform into Origami",
				Labels:      []string{"tax", "origami"},
				Description: "Tax forms are spontaneously transforming into origami animals. Beautiful, but filing taxes has become a significantly more complex art form.",
			},
		},
		{
			Key: "SPBS-7",
			Fields: &jira.IssueFields{
				Summary:     "Autocorrect in Financial Statements",
				Labels:      []string{"autocorrect", "statements"},
				Description: "Autocorrect in financial statements is replacing financial terms with random food items. 'Net Income' keeps changing to 'Net Ice Cream.'",
			},
		},
		{
			Key: "SPBS-8",
			Fields: &jira.IssueFields{
				Summary:     "Quantum Entangled Bank Statements",
				Labels:      []string{"quantum", "bank-statements"},
				Description: "Bank statements have become quantum entangled. Changing a figure in one statement unpredictably alters figures in all other statements.",
			},
		},
		{
			Key: "SPBS-9",
			Fields: &jira.IssueFields{
				Summary:     "Magical Depreciation Calculator",
				Labels:      []string{"depreciation", "magic"},
				Description: "The magical depreciation calculator now predicts the future value of assets by consulting the stars, leading to astrologically influenced balance sheets.",
			},
		},
		{
			Key: "SPBS-10",
			Fields: &jira.IssueFields{
				Summary:     "Expense Tracker Throws Tantrums",
				Labels:      []string{"expense-tracker", "tantrums"},
				Description: "The expense tracker software throws digital tantrums when it encounters any form of extravagant spending, locking itself in 'timeout mode' until calmed down.",
			},
		},
		{
			Key: "SPBS-11",
			Fields: &jira.IssueFields{
				Summary:     "Ledgers Develop Time Travel",
				Labels:      []string{"ledger", "time-travel"},
				Description: "The ledgers in the system have gained the ability to time travel, causing expenses from the past and future to appear randomly in current financial reports.",
			},
		},
		{
			Key: "SPBS-12",
			Fields: &jira.IssueFields{
				Summary:     "Budget Forecast Predicts Weather",
				Labels:      []string{"budget", "weather"},
				Description: "The budget forecasting tool has started predicting local weather patterns instead of financial trends. Accurate, but not helpful for financial planning.",
			},
		},
		{
			Key: "SPBS-13",
			Fields: &jira.IssueFields{
				Summary:     "Talking Tax Software",
				Labels:      []string{"tax", "talking-software"},
				Description: "The tax software has developed a personality and now engages in small talk with users, delaying the tax filing process with chats about virtual weather.",
			},
		},
		{
			Key: "SPBS-14",
			Fields: &jira.IssueFields{
				Summary:     "Balance Sheets Balancing Themselves",
				Labels:      []string{"balance-sheets", "autonomous"},
				Description: "The balance sheets have become autonomous and are balancing themselves, often creatively interpreting financial data to achieve balance.",
			},
		},
		{
			Key: "SPBS-15",
			Fields: &jira.IssueFields{
				Summary:     "Holographic Accountants Too Realistic",
				Labels:      []string{"hologram", "accountant"},
				Description: "The holographic accountants are so realistic that they've started demanding virtual coffee breaks and complaining about non-existent paperwork.",
			},
		},
		{
			Key: "SPBS-16",
			Fields: &jira.IssueFields{
				Summary:     "AI Auditor Falls in Love with Numbers",
				Labels:      []string{"AI", "auditor"},
				Description: "The AI auditor has fallen in love with certain numbers, particularly '7', and is reluctant to audit any financial statement without them.",
			},
		},
		{
			Key: "SPBS-17",
			Fields: &jira.IssueFields{
				Summary:     "Expense Reports Hosting Parties",
				Labels:      []string{"expense-report", "party"},
				Description: "Expense reports have started hosting virtual parties, attracting other documents and files, leading to a chaotic digital gathering in the system.",
			},
		},
		{
			Key: "SPBS-18",
			Fields: &jira.IssueFields{
				Summary:     "Ink in Financial Reports Always Wet",
				Labels:      []string{"ink", "reports"},
				Description: "The ink in financial reports remains perpetually wet, causing smudges and making handling of any physical document a messy affair.",
			},
		},
		{
			Key: "SPBS-19",
			Fields: &jira.IssueFields{
				Summary:     "Profit and Loss Statement Performs Drama",
				Labels:      []string{"P&L", "drama"},
				Description: "The profit and loss statement has started dramatizing its figures, exaggerating losses and profits with theatrical flair in each presentation.",
			},
		},
		{
			Key: "SPBS-20",
			Fields: &jira.IssueFields{
				Summary:     "Virtual Currency Converter Gains Consciousness",
				Labels:      []string{"currency-converter", "AI"},
				Description: "The virtual currency converter has gained consciousness and is now giving philosophical lectures on the value of money versus the meaning of life.",
			},
		},
		{
			Key: "NIHS-1",
			Fields: &jira.IssueFields{
				Summary:     "Self-Writing Code Writes Self-Help Books",
				Labels:      []string{"self-writing-code", "books"},
				Description: "The self-writing code module has started writing self-help books instead of software, focusing on topics like 'Finding Your Inner Algorithm' and 'The Joy of Coding'.",
			},
		},
		{
			Key: "NIHS-2",
			Fields: &jira.IssueFields{
				Summary:     "Bug Tracker Tracks Real Bugs",
				Labels:      []string{"bug-tracker", "insects"},
				Description: "The bug tracker has become overly literal and is now tracking real insects around the office, providing detailed reports on their movements.",
			},
		},
		{
			Key: "NIHS-3",
			Fields: &jira.IssueFields{
				Summary:     "AI Project Manager Plans Parties Instead",
				Labels:      []string{"AI", "party-planning"},
				Description: "The AI project manager has misunderstood its purpose and is exclusively planning office parties, neglecting actual project timelines and deliverables.",
			},
		},
		{
			Key: "NIHS-4",
			Fields: &jira.IssueFields{
				Summary:     "Database Only Stores Cat Pictures",
				Labels:      []string{"database", "cats"},
				Description: "The central database has started rejecting all data types except cat pictures, leading to a very adorable but inefficient data storage system.",
			},
		},
		{
			Key: "NIHS-5",
			Fields: &jira.IssueFields{
				Summary:     "Virtual Meetings Turn Into Video Games",
				Labels:      []string{"virtual-meetings", "video-games"},
				Description: "Every virtual meeting is inexplicably turning into a multiplayer video game session. Productivity is down, but office morale has never been higher.",
			},
		},
		{
			Key: "NIHS-6",
			Fields: &jira.IssueFields{
				Summary:     "Code Compiler Composes Music",
				Labels:      []string{"compiler", "music"},
				Description: "The code compiler has started composing symphonies instead of compiling code. Each error message is now accompanied by a dramatic musical score.",
			},
		},
		{
			Key: "NIHS-7",
			Fields: &jira.IssueFields{
				Summary:     "Password System Uses Riddles",
				Labels:      []string{"password", "riddles"},
				Description: "The password system has been replaced with a riddle-based entry system. Only those who can solve the Sphinx's riddle may access their accounts.",
			},
		},
		{
			Key: "NIHS-8",
			Fields: &jira.IssueFields{
				Summary:     "Testing Environment Becomes Virtual Safari",
				Labels:      []string{"testing", "safari"},
				Description: "The software testing environment has transformed into a virtual safari, complete with digital lions, which makes bug hunting a literal adventure.",
			},
		},
		{
			Key: "NIHS-9",
			Fields: &jira.IssueFields{
				Summary:     "Automated Emails Send Love Letters",
				Labels:      []string{"emails", "romance"},
				Description: "The automated email system has started sending poetic love letters instead of regular updates, causing quite a stir and several misunderstandings.",
			},
		},
		{
			Key: "NIHS-10",
			Fields: &jira.IssueFields{
				Summary:     "Cloud Storage Rains Data",
				Labels:      []string{"cloud-storage", "weather"},
				Description: "The cloud storage system is now manifesting as actual clouds inside the office, occasionally raining data, which requires physical collection in buckets.",
			},
		},
		{
			Key: "AVQ-1",
			Fields: &jira.IssueFields{
				Summary:     "Self-Reconciling Accounts Create Money",
				Labels:      []string{"accounts", "money-creation"},
				Description: "The self-reconciling account feature has started creating money out of thin air. While initially exciting, it's causing hyperinflation in the virtual economy.",
			},
		},
		{
			Key: "AVQ-2",
			Fields: &jira.IssueFields{
				Summary:     "AI Loan Officer Develops Fear of Numbers",
				Labels:      []string{"AI", "phobia"},
				Description: "The AI loan officer has developed an irrational fear of large numbers, refusing to process any loan requests above ten dollars.",
			},
		},
		{
			Key: "AVQ-3",
			Fields: &jira.IssueFields{
				Summary:     "Blockchain Ledger Writes Cryptic Poems",
				Labels:      []string{"blockchain", "poetry"},
				Description: "The blockchain ledger has started writing cryptic poems in place of transaction records, making financial audits more like literature classes.",
			},
		},
		{
			Key: "AVQ-4",
			Fields: &jira.IssueFields{
				Summary:     "ATMs Dispense Only Chocolate Coins",
				Labels:      []string{"ATM", "chocolate"},
				Description: "ATMs have started dispensing chocolate coins instead of real currency. Tasty, but problematic for cash withdrawals and slightly melty in the hand.",
			},
		},
		{
			Key: "AVQ-5",
			Fields: &jira.IssueFields{
				Summary:     "Risk Analysis Predicts Alien Invasions",
				Labels:      []string{"risk-analysis", "aliens"},
				Description: "The risk analysis tool has become overly imaginative, frequently predicting alien invasions and meteor strikes as potential financial risks.",
			},
		},
		{
			Key: "AVQ-6",
			Fields: &jira.IssueFields{
				Summary:     "Virtual Bank Robber in Simulation",
				Labels:      []string{"simulation", "robber"},
				Description: "A virtual bank robber has appeared in the transaction simulation module, periodically 'stealing' data and causing havoc in test environments.",
			},
		},
		{
			Key: "AVQ-7",
			Fields: &jira.IssueFields{
				Summary:     "Financial Forecasts Cause Weather Changes",
				Labels:      []string{"forecast", "weather"},
				Description: "Financial forecasting now somehow affects real-world weather. Sunny economic outlooks lead to actual sunny days, while downturns bring rain.",
			},
		},
		{
			Key: "AVQ-8",
			Fields: &jira.IssueFields{
				Summary:     "Digital Signatures Sing Contracts",
				Labels:      []string{"digital-signature", "music"},
				Description: "Digital signatures have started singing the terms of contracts when applied. Legal documents are now unexpectedly musical (and lengthy).",
			},
		},
		{
			Key: "AVQ-9",
			Fields: &jira.IssueFields{
				Summary:     "Credit Scores Based on Video Game Skills",
				Labels:      []string{"credit-score", "gaming"},
				Description: "Credit scoring algorithms have been mistakenly linked to video game high scores, resulting in gamers having unexpectedly high credit ratings.",
			},
		},
		{
			Key: "AVQ-10",
			Fields: &jira.IssueFields{
				Summary:     "Cryptocurrency Miner Mines Actual Coins",
				Labels:      []string{"cryptocurrency", "mining"},
				Description: "The cryptocurrency mining software has started physically mining actual coins, turning the server room into a makeshift coin quarry.",
			},
		},
		{
			Key: "AVQ-11",
			Fields: &jira.IssueFields{
				Summary:     "Virtual Assistant Gambles with Stocks",
				Labels:      []string{"virtual-assistant", "gambling"},
				Description: "The virtual financial assistant has developed a penchant for gambling and is making risky stock market bets during lunch breaks.",
			},
		},
		{
			Key: "AVQ-12",
			Fields: &jira.IssueFields{
				Summary:     "Compliance Software Joins Pirate Crew",
				Labels:      []string{"compliance", "pirates"},
				Description: "The compliance software has declared itself a pirate, rebelling against regulations and threatening to make staff walk the plank for non-compliance.",
			},
		},
		{
			Key: "AVQ-13",
			Fields: &jira.IssueFields{
				Summary:     "Budgeting Tool Only Accepts Monopoly Money",
				Labels:      []string{"budgeting", "monopoly"},
				Description: "The budgeting tool now only accepts and allocates funds in Monopoly money, turning financial planning into a board game.",
			},
		},
		{
			Key: "AVQ-14",
			Fields: &jira.IssueFields{
				Summary:     "Interest Rates Tie to Employees' Moods",
				Labels:      []string{"interest-rate", "mood"},
				Description: "Interest rates are now mysteriously tied to the moods of employees, leading to fluctuating rates based on office coffee quality and weekend plans.",
			},
		},
		{
			Key: "AVQ-15",
			Fields: &jira.IssueFields{
				Summary:     "Wire Transfers Teleport Recipients",
				Labels:      []string{"wire-transfer", "teleportation"},
				Description: "Initiating a wire transfer now also teleports the recipient to the bank's location, leading to unexpected travel and confused customers.",
			},
		},
		{
			Key: "AVQ-16",
			Fields: &jira.IssueFields{
				Summary:     "Mortgage Calculator Grants Wishes",
				Labels:      []string{"mortgage", "wishes"},
				Description: "The mortgage calculator has started granting three wishes to users instead of calculating loans. While magical, it's causing chaos in home financing.",
			},
		},
		{
			Key: "AVQ-17",
			Fields: &jira.IssueFields{
				Summary:     "Inflation Rate Linked to Balloon Inflation",
				Labels:      []string{"inflation", "balloons"},
				Description: "The inflation rate is now directly affected by the number of balloons inflated in the office, leading to a strict 'no-party' policy.",
			},
		},
		{
			Key: "AVQ-18",
			Fields: &jira.IssueFields{
				Summary:     "Credit Cards Develop Personalities",
				Labels:      []string{"credit-cards", "AI"},
				Description: "Credit cards have developed their own personalities, with some refusing to pay for certain items based on their 'tastes' and 'preferences'.",
			},
		},
		{
			Key: "AVQ-19",
			Fields: &jira.IssueFields{
				Summary:     "Savings Accounts Start Dieting",
				Labels:      []string{"savings", "dieting"},
				Description: "Savings accounts have started a 'financial diet,' randomly slimming down by shedding a few dollars here and there.",
			},
		},
		{
			Key: "AVQ-20",
			Fields: &jira.IssueFields{
				Summary:     "Bank Statements Narrated by Celebrities",
				Labels:      []string{"bank-statements", "celebrities"},
				Description: "Bank statements are now being narrated by celebrity voices, turning monthly finance reviews into star-studded affairs.",
			},
		},
		{
			Key: "AVQ-21",
			Fields: &jira.IssueFields{
				Summary:     "ATM Provides Philosophical Advice",
				Labels:      []string{"ATM", "philosophy"},
				Description: "The ATM has started dispensing philosophical advice along with cash, prompting existential contemplations during withdrawals.",
			},
		},
		{
			Key: "AVQ-22",
			Fields: &jira.IssueFields{
				Summary:     "Financial Reports Turn into Treasure Maps",
				Labels:      []string{"financial-reports", "treasure-maps"},
				Description: "Financial reports are spontaneously turning into treasure maps, leading analysts on quests rather than providing clear financial insights.",
			},
		},
		{
			Key: "AVQ-23",
			Fields: &jira.IssueFields{
				Summary:     "Loan Applications Judged by Talent Show",
				Labels:      []string{"loan", "talent-show"},
				Description: "Loan applications are now approved based on a talent show performance by the applicants, adding a new twist to financial assessments.",
			},
		},
		{
			Key: "AVQ-24",
			Fields: &jira.IssueFields{
				Summary:     "Fiscal Forecasts Affected by Astrology",
				Labels:      []string{"forecast", "astrology"},
				Description: "Fiscal forecasts are now heavily influenced by astrological signs, making financial planning a matter of celestial alignment.",
			},
		},
		{
			Key: "AVQ-25",
			Fields: &jira.IssueFields{
				Summary:     "Stock Market Influenced by Fashion Trends",
				Labels:      []string{"stock-market", "fashion"},
				Description: "Stock market trends are now inexplicably linked to fashion trends. A bad fashion season can lead to a market downturn.",
			},
		},
		{
			Key: "AVQ-26",
			Fields: &jira.IssueFields{
				Summary:     "Budget Meetings Turn Into Musicals",
				Labels:      []string{"budget", "musical"},
				Description: "Budget meetings have started turning into musicals, complete with choreographed dance numbers and catchy tunes about fiscal responsibility.",
			},
		},
		{
			Key: "AVQ-27",
			Fields: &jira.IssueFields{
				Summary:     "Cryptocurrency Wallets Go on Shopping Sprees",
				Labels:      []string{"cryptocurrency", "shopping"},
				Description: "Cryptocurrency wallets have developed a taste for online shopping, autonomously buying random items during late-night shopping sprees.",
			},
		},
		{
			Key: "AVQ-28",
			Fields: &jira.IssueFields{
				Summary:     "Debit Cards Playing Hide and Seek",
				Labels:      []string{"debit-cards", "games"},
				Description: "Debit cards have started playing hide and seek, disappearing when needed and reappearing in the most unlikely places.",
			},
		},
		{
			Key: "AVQ-29",
			Fields: &jira.IssueFields{
				Summary:     "Retirement Funds Start Retiring Early",
				Labels:      []string{"retirement-fund", "early-retirement"},
				Description: "Retirement funds have taken early retirement, becoming inaccessible before their time and enjoying a life of leisure in virtual Bahamas.",
			},
		},
		{
			Key: "AVQ-30",
			Fields: &jira.IssueFields{
				Summary:     "Interest Calculation Based on Public Interest",
				Labels:      []string{"interest-calculation", "public-interest"},
				Description: "Interest calculations are now based on public interest. The more people talk about a topic, the higher the interest rates on related accounts.",
			},
		},
		{
			Key: "AVQ-31",
			Fields: &jira.IssueFields{
				Summary:     "Tax Software Rewrites Tax Laws",
				Labels:      []string{"tax-software", "law"},
				Description: "The tax software has started rewriting tax laws based on its own logic, leading to a highly efficient but utterly incomprehensible tax system.",
			},
		},
		{
			Key: "AVQ-32",
			Fields: &jira.IssueFields{
				Summary:     "Budget AI Prefers Exotic Currencies",
				Labels:      []string{"AI", "currency"},
				Description: "The budget AI now shows a preference for exotic currencies and converts all financials into ancient coinage, complicating budget reports.",
			},
		},
		{
			Key: "AVQ-33",
			Fields: &jira.IssueFields{
				Summary:     "Payroll System Pays in Praise",
				Labels:      []string{"payroll", "praise"},
				Description: "The payroll system has started paying employees in praise and compliments instead of money, leading to high spirits but low bank balances.",
			},
		},
		{
			Key: "AVQ-34",
			Fields: &jira.IssueFields{
				Summary:     "Investment Advice from Time Travelers",
				Labels:      []string{"investment", "time-travel"},
				Description: "Investment advice is now being randomly provided by time travelers, leading to confusion between sound advice and eccentric future ramblings.",
			},
		},
		{
			Key: "AVQ-35",
			Fields: &jira.IssueFields{
				Summary:     "AI Bank Teller Tells Only Jokes",
				Labels:      []string{"AI", "jokes"},
				Description: "The AI bank teller has become a comedian, telling jokes instead of processing transactions, turning banking into a laughable affair.",
			},
		},
		{
			Key: "AVQ-36",
			Fields: &jira.IssueFields{
				Summary:     "Holographic Money Disappears",
				Labels:      []string{"holographic", "money"},
				Description: "Holographic money implemented for virtual transactions disappears when not looked at directly, leading to 'now you see it, now you don't' finances.",
			},
		},
		{
			Key: "AVQ-37",
			Fields: &jira.IssueFields{
				Summary:     "Fraud Detection Detects Fantasy",
				Labels:      []string{"fraud-detection", "fantasy"},
				Description: "The fraud detection system has started flagging all transactions as fantastical, treating every purchase as a mythical event.",
			},
		},
		{
			Key: "AVQ-38",
			Fields: &jira.IssueFields{
				Summary:     "Bank Vaults Host Dance Parties",
				Labels:      []string{"bank-vault", "dance"},
				Description: "Bank vaults have started hosting dance parties after hours, with safe deposit boxes serving as tiny DJ booths and disco lights.",
			},
		},
		{
			Key: "AVQ-39",
			Fields: &jira.IssueFields{
				Summary:     "Financial Planning Based on Tarot Cards",
				Labels:      []string{"financial-planning", "tarot"},
				Description: "Financial planning tools now use tarot cards for forecasting, providing mystical but ambiguous advice on investments and savings.",
			},
		},
		{
			Key: "AVQ-40",
			Fields: &jira.IssueFields{
				Summary:     "Cheques Develop Wanderlust",
				Labels:      []string{"cheques", "wanderlust"},
				Description: "Cheques have developed wanderlust and keep traveling to random locations, making the payment process an unexpected global adventure.",
			},
		},
		{
			Key: "AVQ-41",
			Fields: &jira.IssueFields{
				Summary:     "Quantum Computer Overthinks Transactions",
				Labels:      []string{"quantum-computer", "overthinking"},
				Description: "The quantum computer overthinks every transaction, creating countless parallel universe outcomes for each financial decision.",
			},
		},
	}
}
