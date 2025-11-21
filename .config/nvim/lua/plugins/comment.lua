return {
  "nvim-mini/mini.comment",
  event = "VeryLazy",
  opts = {
    mappings = {
      -- Toggle comment (like `gcip` - comment inner paragraph) for both
      -- Normal and Visual modes
      comment = "gc",

      -- Toggle comment on current line
      comment_line = "<leader>c",

      -- Toggle comment on visual selection
      comment_visual = "<leader>c",

      -- Define 'comment' textobject (like `dgc` - delete whole comment block)
      textobject = "gc",
    },
  },
}
