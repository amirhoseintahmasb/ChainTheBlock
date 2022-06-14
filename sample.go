package main
//find the calldata for swap
func SwapOpts(ctx context.Context, network, from, to string, amount, gasLimit string) (string, error) {
	client := go1inch.NewClient()
	res, _, err := client.Swap(ctx, network, from, to, amount, gasLimit)
	if err != nil {
		return "", err
	}
	return res, nil
}