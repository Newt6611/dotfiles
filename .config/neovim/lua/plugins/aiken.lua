return {
	"aiken-lang/editor-integration-nvim",
	dependencies = {
		'neovim/nvim-lspconfig',
	},
	init = function()
		require'lspconfig'.aiken.setup({})
	end,
}
