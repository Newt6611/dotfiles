vim.pack.add({
  { src = 'https://github.com/NeogitOrg/neogit' },
  { src = 'https://github.com/nvim-lua/plenary.nvim' },
  { src = 'https://github.com/sindrets/diffview.nvim' },
})

local neogit = require('neogit')
neogit.setup({
  floating = {
    relative = "editor",
    width = 0.9,
    height = 0.9,
    style = "minimal",
    border = "rounded",
  },
})

local map = vim.keymap.set
map('n', '<leader>gh', function() neogit.open({ kind = "floating" }) end, { desc = 'Open Neogit' })
