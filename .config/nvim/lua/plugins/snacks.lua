vim.pack.add({
  { src = 'https://github.com/folke/snacks.nvim' },
})

require('snacks').setup({
  picker = {
    enabled = true,
  },
  indent = {
    enabled = true,
    animate = {
      enabled = false,
    },
  },
  lazygit = {
    enabled = true,
  },
})

local map = vim.keymap.set
map('n', "<leader>pf", function() Snacks.picker.files() end, { desc = "Find Files" })
map('n', "<leader>ps", function() Snacks.picker.grep() end, { desc = "Grep" })
map('n', "<leader>gd", function() Snacks.picker.git_diff() end, { desc = "Git Diff (Hunks)" })

map('n', "<leader>gl", function() Snacks.lazygit.log() end, { desc = "LazyGit Log" })
map('n', "<leader>gg", function() Snacks.lazygit() end, { desc = "LazyGit" })
