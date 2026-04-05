return {
  { 'rafamadriz/friendly-snippets' },
  {
    'saghen/blink.cmp',
    version = '1.*',
    build = 'cargo build --release',
    opts = {
      keymap = {
        ['<Up>'] = { 'select_prev', 'fallback' },
        ['<Down>'] = { 'select_next', 'fallback' },
        ['<CR>'] = { 'accept', 'fallback' },
      },
      appearance = {
        nerd_font_variant = 'mono'
      },
      completion = { documentation = { auto_show = false } },
      sources = {
        default = { 'lsp', 'path', 'snippets', 'buffer' },
      },
      fuzzy = { implementation = 'prefer_rust' }
    },
    opts_extend = { "sources.default" }
  }
}
