vim.pack.add({
  { src = 'https://github.com/github/copilot.vim' }
})

local map = vim.keymap.set
local g = vim.g

map('i', '<C-J>', 'copilot#Accept("\\<CR>")', {
  expr = true,
  replace_keycodes = false
})

g.copilot_no_tab_map = true
