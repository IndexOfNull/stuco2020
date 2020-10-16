module.exports = {
  future: {
    removeDeprecatedGapUtilities: true,
    purgeLayersByDefault: true,
    defaultLineHeights: true,
    standardFontWeights: true
  },
  purge: ["./src/**/*.vue", "./src/**/*.html"],
  theme: {
    extend: {
      colors: {
        danger: "#D7263D",
        dark: "#02182B",
        primary: "#0197F6",
        secondary: "#68C5DB",
        tertiary: "#448FA3"
      }
    }
  },
  variants: {},
  plugins: []
};
