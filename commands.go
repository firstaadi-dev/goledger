package goledger

import "context"

// Accounts runs `hledger accounts`.
func (c *Client) Accounts(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "accounts", req)
}

// Activity runs `hledger activity`.
func (c *Client) Activity(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "activity", req)
}

// Add runs `hledger add`.
func (c *Client) Add(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "add", req)
}

// ARegister runs `hledger aregister`.
func (c *Client) ARegister(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "aregister", req)
}

// Balance runs `hledger balance`.
func (c *Client) Balance(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "balance", req)
}

// BalanceSheet runs `hledger balancesheet`.
func (c *Client) BalanceSheet(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "balancesheet", req)
}

// BS runs `hledger bs` alias.
func (c *Client) BS(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "bs", req)
}

// Cashflow runs `hledger cashflow`.
func (c *Client) Cashflow(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "cashflow", req)
}

// CF runs `hledger cf` alias.
func (c *Client) CF(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "cf", req)
}

// Check runs `hledger check`.
func (c *Client) Check(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "check", req)
}

// Close runs `hledger close`.
func (c *Client) Close(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "close", req)
}

// Codes runs `hledger codes`.
func (c *Client) Codes(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "codes", req)
}

// Commodities runs `hledger commodities`.
func (c *Client) Commodities(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "commodities", req)
}

// Descriptions runs `hledger descriptions`.
func (c *Client) Descriptions(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "descriptions", req)
}

// Diff runs `hledger diff`.
func (c *Client) Diff(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "diff", req)
}

// Files runs `hledger files`.
func (c *Client) Files(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "files", req)
}

// Help runs `hledger help`.
func (c *Client) Help(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "help", req)
}

// Import runs `hledger import`.
func (c *Client) Import(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "import", req)
}

// IncomeStatement runs `hledger incomestatement`.
func (c *Client) IncomeStatement(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "incomestatement", req)
}

// IS runs `hledger is` alias.
func (c *Client) IS(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "is", req)
}

// Notes runs `hledger notes`.
func (c *Client) Notes(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "notes", req)
}

// Payees runs `hledger payees`.
func (c *Client) Payees(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "payees", req)
}

// Prices runs `hledger prices`.
func (c *Client) Prices(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "prices", req)
}

// Print runs `hledger print`.
func (c *Client) Print(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "print", req)
}

// Register runs `hledger register`.
func (c *Client) Register(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "register", req)
}

// Rewrite runs `hledger rewrite`.
func (c *Client) Rewrite(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "rewrite", req)
}

// ROI runs `hledger roi`.
func (c *Client) ROI(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "roi", req)
}

// Stats runs `hledger stats`.
func (c *Client) Stats(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "stats", req)
}

// Tags runs `hledger tags`.
func (c *Client) Tags(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "tags", req)
}

// Test runs `hledger test`.
func (c *Client) Test(ctx context.Context, req CommandRequest) (CommandResult, error) {
	return c.runCommand(ctx, "test", req)
}
