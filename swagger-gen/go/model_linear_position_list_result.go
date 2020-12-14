/*
 * Bybit API
 *
 * ## REST API for the Bybit Exchange. Base URI: [https://api.bybit.com]  
 *
 * API version: 0.2.10
 * Contact: support@bybit.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type LinearPositionListResult struct {
	BustPrice float64 `json:"bust_price,omitempty"`
	CumRealisedPnl float64 `json:"cum_realised_pnl,omitempty"`
	EntryPrice float64 `json:"entry_price,omitempty"`
	FreeQty float64 `json:"free_qty,omitempty"`
	Leverage float64 `json:"leverage,omitempty"`
	LiqPrice float64 `json:"liq_price,omitempty"`
	OccClosingFee float64 `json:"occ_closing_fee,omitempty"`
	PositionMargin float64 `json:"position_margin,omitempty"`
	PositionValue float64 `json:"position_value,omitempty"`
	RealisedPnl float64 `json:"realised_pnl,omitempty"`
	Side string `json:"side,omitempty"`
	Size float64 `json:"size,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	UserId int64 `json:"user_id,omitempty"`
	TpSlMode string `json:"tp_sl_mode,omitempty"`
}
