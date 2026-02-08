vim.pack.add({
	{ src = 'https://github.com/lewis6991/gitsigns.nvim' }
})

require('gitsigns').setup()

local map = vim.keymap.set

map({ "n", "v" }, "<leader>r", ":Gitsigns reset_hunk<CR>", "Reset Hunk")
