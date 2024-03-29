// Code generated by entc, DO NOT EDIT.

package ent

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
}

// hooks per client, for fast access.
type hooks struct {
	Agentkyc                   []ent.Hook
	Agenttype                  []ent.Hook
	Areahistory                []ent.Hook
	Bankdetail                 []ent.Hook
	Bulk                       []ent.Hook
	Configarea                 []ent.Hook
	Configdatecalculaterank    []ent.Hook
	Configoccupation           []ent.Hook
	Configpoint                []ent.Hook
	Consumer                   []ent.Hook
	Fileimport                 []ent.Hook
	Fileinsert                 []ent.Hook
	Limitpoint                 []ent.Hook
	Loanbinding                []ent.Hook
	Logexport                  []ent.Hook
	MerchantTransaction        []ent.Hook
	Occupationhistory          []ent.Hook
	Pendingkyc                 []ent.Hook
	Pendingloanbinding         []ent.Hook
	Pointcsv                   []ent.Hook
	Pointkycrv                 []ent.Hook
	Pointpendingkyctransaction []ent.Hook
	Pointpendinglbtransaction  []ent.Hook
	Pointtransaction           []ent.Hook
	Ranking                    []ent.Hook
	ReportWallet               []ent.Hook
	Reportwallettb             []ent.Hook
	StatementEndingBalanc      []ent.Hook
	Transactionfactor          []ent.Hook
	Transactionfactorhistory   []ent.Hook
	Transactionfactoritem      []ent.Hook
	Transactionfactoritemtmp   []ent.Hook
	Userprofile                []ent.Hook
	Userwallet                 []ent.Hook
	Watchlist                  []ent.Hook
	Watchlisthistory           []ent.Hook
	Watchlisttype              []ent.Hook
	Writelog                   []ent.Hook
}

// Options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...interface{})) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}
