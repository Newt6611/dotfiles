vim.pack.add({
  { src = 'https://github.com/rachartier/tiny-inline-diagnostic.nvim' }
})

require("tiny-inline-diagnostic").setup()

-- Disable Neovim's default virtual text diagnostics
vim.diagnostic.config({ virtual_text = false })
