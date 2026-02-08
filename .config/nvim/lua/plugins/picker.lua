vim.pack.add({
  { src = 'https://github.com/nvim-mini/mini.pick', checkout = 'stable' }
})

local pick = require('mini.pick')
pick.setup()

local map = vim.keymap.set
map('n', '<leader>pf', pick.builtin.files)

map('n', '<leader>ps', pick.builtin.grep)
