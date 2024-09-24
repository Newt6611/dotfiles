return {
	"lewis6991/gitsigns.nvim",
	event = "LazyFile",
	opts = {
		signs = {
			add = { text = "▎" },
			change = { text = "▎" },
			delete = { text = "" },
			topdelete = { text = "" },
			changedelete = { text = "▎" },
			untracked = { text = "▎" },
		},
		on_attach = function(buffer)
			local gs = package.loaded.gitsigns

			local function map(mode, l, r, desc)
				vim.keymap.set(mode, l, r, { buffer = buffer, desc = desc })
			end

			-- stylua: ignore start
			map("n", "<leader>[", gs.next_hunk, "Next Hunk")
			map("n", "<leader>]", gs.prev_hunk, "Prev Hunk")

			map({ "n", "v" }, "<leader>s", ":Gitsigns stage_hunk<CR>", "Stage Hunk")
			map({ "n", "v" }, "<leader>r", ":Gitsigns reset_hunk<CR>", "Reset Hunk")
			map('n', '<leader>hp', '<cmd>Gitsigns preview_hunk<CR>')
			map("n", "<leader>hu", gs.undo_stage_hunk, "Undo Stage Hunk")

			map("n", "<leader>ghS", gs.stage_buffer, "Stage Buffer")
			map("n", "<leader>ghR", gs.reset_buffer, "Reset Buffer")
			map("n", "<leader>ghp", gs.preview_hunk_inline, "Preview Hunk Inline")
			map("n", "<leader>b", function() gs.blame_line({ full = false }) end, "Blame Line")
			map("n", "<leader>ghd", gs.diffthis, "Diff This")
			map("n", "<leader>ghD", function() gs.diffthis("~") end, "Diff This ~")
			map({ "o", "x" }, "ih", ":<C-U>Gitsigns select_hunk<CR>", "GitSigns Select Hunk")
		end,
	},
}
