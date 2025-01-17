package sarif

type Tool struct {
	Driver *Driver `json:"driver"`
}

type Driver struct {
	Name           string  `json:"name"`
	InformationURI string  `json:"informationUri"`
	Rules          []*Rule `json:"rules,omitempty"`
}

// Rule specifies a Sarif Rule object
type Rule struct {
	ID               string            `json:"id"`
	ShortDescription *TextBlock        `json:"shortDescription"`
	HelpURI          string            `json:"helpUri,omitempty"`
	Help             *TextBlock        `json:"help,omitempty"`
	Properties       map[string]string `json:"properties,omitempty"`
}

func (driver *Driver) getOrCreateRule(rule *Rule) uint {
	for i, r := range driver.Rules {
		if r.ID == rule.ID {
			return uint(i)
		}
	}
	driver.Rules = append(driver.Rules, rule)
	return uint(len(driver.Rules) - 1)
}

func newRule(ruleID string) *Rule {
	return &Rule{
		ID: ruleID,
	}
}

// WithDescription specifies a description for a rule and returns the updated rule
func (rule *Rule) WithDescription(description string) *Rule {
	rule.ShortDescription = &TextBlock{
		Text: description,
	}
	return rule
}

// WithHelpURI specifies a helpURI for a rule and returns the updated rule
func (rule *Rule) WithHelpURI(helpURI string) *Rule {
	rule.HelpURI = helpURI
	return rule
}

// WithHelp specifies a help text  for a rule and returns the updated rule
func (rule *Rule) WithHelp(helpText string) *Rule {
	rule.Help = &TextBlock{
		Text: helpText,
	}
	return rule
}

// WithProperties specifies properties for a rule and returns the updated rule
func (rule *Rule) WithProperties(properties map[string]string) *Rule {
	rule.Properties = properties
	return rule
}
