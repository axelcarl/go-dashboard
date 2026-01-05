export const paths = {
  home: {
    path: "/",
    getHref: () => "/",
  },
  app: {
    root: {
      path: "/app",
      getHref: () => "/app",
    },
    dashboard: {
      path: "",
      getHref: () => "/app",
    },
  },
} as const;
