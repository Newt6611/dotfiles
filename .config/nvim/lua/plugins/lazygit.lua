vim.pack.add({
  { src = 'https://github.com/kdheepak/lazygit.nvim' }
})

local map = vim.keymap.set

map('n', "<leader>gg", ":LazyGit<CR>", { desc = "Lazygit" })
