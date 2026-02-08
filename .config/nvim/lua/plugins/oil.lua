vim.pack.add({
  { src = 'https://github.com/stevearc/oil.nvim' }
})

require("oil").setup({
  use_default_keymaps = false,
  keymaps = {
    ["<CR>"] = "actions.select",
    ["-"] = { "actions.parent", mode = "n" },
    ["_"] = { "actions.open_cwd", mode = "n" },
  },
  view_options = {
    show_hidden = true,
  },
  float = {
    border = 'rounded',
  }
})

local map = vim.keymap.set
map("n", "-", "<CMD>Oil --float<CR>", { desc = "Open parent directory" })
