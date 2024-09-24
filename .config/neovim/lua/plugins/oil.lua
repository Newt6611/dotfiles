return {
	'stevearc/oil.nvim',
	lazy = false,
	opts = {
		view_options = {
			show_hidden = true
		}
	},

	-- Optional dependencies
	dependencies = { "nvim-tree/nvim-web-devicons" },

	-- ["-"] = "actions.parent",
	-- ["_"] = "actions.open_cwd",
	keys = function()
		return {
			{ "-", function ()
				require("oil").toggle_float()
			end},

			{ "<space>-", "<CMD>Oil<CR>", { desc = "Open parent directory" }}
		}
	end
}
