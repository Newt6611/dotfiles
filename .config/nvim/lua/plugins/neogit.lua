vim.pack.add({
  { src = 'https://github.com/NeogitOrg/neogit' }
})

require('neogit').setup({
  graph_style = "unicode",
  floating = {
    relative = "editor",
    width = 0.9,
    height = 0.9,
    style = "minimal",
    border = "rounded",
  },
})

local map = vim.keymap.set
map('n', '<leader>gh', '<CMD>Neogit kind=floating<CR>', { desc = 'Open Neogit UI' })
