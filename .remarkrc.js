import stringWidth from "string-width"

export default {
  plugins: [
    "preset-lint-consistent",
    "preset-lint-recommended",
    ["gfm", { stringLength: c => stringWidth(c, { ambiguousIsNarrow: false }) }],
  ]
}
