vim.pack.add({
  { src = 'https://github.com/folke/tokyonight.nvim' }
})

require('tokyonight').setup({
  style = "storm",
  transparent = true,
  styles = {
    sidebars = "transparent",
    floats = "transparent",
  },
})

vim.cmd[[colorscheme tokyonight]]
