return {
	"nvim-telescope/telescope.nvim",
	dependencies = {
		"nvim-telescope/telescope-fzf-native.nvim",
		build = "make",
	},

	-- replace all Telescope keymaps with only one mapping
	keys = function()
		return {
			{ "<leader>pf", "<cmd>Telescope find_files<cr>" },
			{ "<leader>pb", "<cmd>Telescope buffers<cr>" },
			{
				"<leader>ps",
				function()
					require("telescope.builtin").grep_string({ search = vim.fn.input("Grep > ") })
				end,
			},
		}
	end,
}
