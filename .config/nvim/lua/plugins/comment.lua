vim.pack.add({
	{ src = 'https://github.com/nvim-mini/mini.comment' }
})

require('mini.comment').setup({
	mappings = {
		comment = '<leader>c',
		comment_line = '<leader>c',
		comment_visual = '<leader>c',
		textobject = '<leader>c',
	},
})
