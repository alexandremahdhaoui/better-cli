package parser

// PSpec -
type PSpec struct {
	Pos  int
	PosR []string
}

// FSpec -
type FSpec map[string]PSpec

// OptRule -
type OptRule struct {
	Short FSpec
	Long  FSpec
}

func NewOptRule() OptRule {
	return OptRule{
		Short: make(FSpec),
		Long:  make(FSpec),
	}
}

// RuleSpec specifies:
//  	- pos: 		How many positional expression this rule expects
//  	- posR: 	Specifies the "Regex Type" of each of the pos
//		- srt:		if it expects a sub srt or not
//		- opt:		if it expects OptRules (true/false)
type RuleSpec struct {
	Pos  int
	PosR []string
	Srt  bool
	Opt  bool
}

// Rule specifies :
// 		- spec:		Syntax specification of this rule
// 		- srt:		Specifies a sub SRT
// 		- opt:		if Rule.spec.opt = true, an OptRule must be specified
type Rule struct {
	Spec RuleSpec
	Opt  OptRule
	Srt  SRT
}

// SRT - Syntax-Rule-Tree
type SRT map[string]Rule

func (srt SRT) getRule(s string) Rule {
	return srt[s]
}
