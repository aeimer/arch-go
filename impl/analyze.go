package impl

import (
	"github.com/fdaines/arch-go/config"
	"github.com/fdaines/arch-go/impl/contents"
	"github.com/fdaines/arch-go/impl/cycles"
	"github.com/fdaines/arch-go/impl/dependencies"
	"github.com/fdaines/arch-go/impl/functions"
	"github.com/fdaines/arch-go/impl/model"
	"github.com/fdaines/arch-go/utils/output"
	"github.com/fdaines/arch-go/utils/packages"
)

func CheckArchitecture(config *config.Config, mainPackage string, pkgs []*packages.PackageInfo) *model.Result {
	result := &model.Result{}

	output.Printf("Analyze Module: %s\n", mainPackage)
	result.DependenciesRulesResults = checkDependencies(config.DependenciesRules, mainPackage, pkgs)
	result.ContentsRuleResults = checkContents(config.ContentRules, mainPackage, pkgs)
	result.CyclesRuleResults = checkCycles(config.CyclesRules, mainPackage, pkgs)
	result.FunctionsRulesResults = checkFunctions(config.FunctionsRules, mainPackage, pkgs)

	return result
}

func checkCycles(rules []config.CyclesRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.CyclesRuleResult {
	results := []*model.CyclesRuleResult{}
	for _, rule := range rules {
		results = cycles.CheckRule(results, rule, mainPackage, pkgs)
	}
	return results
}

func checkFunctions(rules []config.FunctionsRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.FunctionsRuleResult {
	results := []*model.FunctionsRuleResult{}
	for _, rule := range rules {
		results = functions.CheckRule(results, rule, mainPackage, pkgs)
	}
	return results
}

func checkDependencies(rules []config.DependenciesRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.DependenciesRuleResult {
	results := []*model.DependenciesRuleResult{}
	for _, r := range rules {
		results = dependencies.CheckDependenciesRule(results, r, mainPackage, pkgs)
	}
	return results
}

func checkContents(rules []config.ContentsRule, mainPackage string, pkgs []*packages.PackageInfo) []*model.ContentsRuleResult {
	results := []*model.ContentsRuleResult{}
	for _, rule := range rules {
		results = contents.CheckRule(results, rule, mainPackage, pkgs)
	}
	return results
}
