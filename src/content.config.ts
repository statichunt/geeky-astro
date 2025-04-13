import { glob } from "astro/loaders";
import { defineCollection, z } from "astro:content";

const aboutCollection = defineCollection({
  loader: glob({ pattern: "**/*.{md,mdx}", base: "src/content/about" }),
  schema: z.object({
    title: z.string().optional(),
    image: z.string().optional(),
    meta_title: z.string().optional(),
    description: z.string().optional(),
    education: z.object({
      title: z.string().optional(),
      degrees: z.array(
        z.object({ university: z.string(), content: z.string() }),
      ),
    }),
    experience: z.object({
      title: z.string().optional(),
      list: z.array(z.string()),
    }),
  }),
});

const contactCollection = defineCollection({
  loader: glob({ pattern: "**/*.{md,mdx}", base: "src/content/contact" }),
  schema: z.object({
    title: z.string().optional(),
    description: z.string().optional(),
    meta_title: z.string().optional(),
    phone: z.string().optional(),
    mail: z.string().optional(),
    location: z.string().optional(),
    form_action: z.string().optional(),
    social_links: z.array(
      z.object({
        icon: z.string(),
        content: z.string(),
        link: z.string().optional(),
      }),
    ),
    draft: z.boolean(),
  }),
});

const homepageCollection = defineCollection({
  loader: glob({ pattern: "**/*.{md,mdx}", base: "src/content/homepage" }),
  schema: z.object({
    banner: z.object({
      title: z.string(),
      title_small: z.string(),
      content: z.string(),
      image_enable: z.boolean(),
      image: z.string(),
      button: z.object({
        enable: z.boolean(),
        label: z.string(),
        link: z.string(),
        rel: z.string(),
      }),
    }),
    featured_posts: z.object({ enable: z.boolean(), title: z.string() }),
    promotion: z.object({
      enable: z.boolean(),
      image: z.string(),
      link: z.string(),
    }),
    recent_posts: z.object({ title: z.string(), enable: z.boolean() }),
  }),
});

const pagesCollection = defineCollection({
  loader: glob({ pattern: "**/*.{md,mdx}", base: "src/content/pages" }),
  schema: z.object({
    title: z.string().optional(),
    meta_title: z.string().optional(),
    description: z.string().optional(),
    draft: z.boolean(),
  }),
});

const postsCollection = defineCollection({
  loader: glob({ pattern: "**/*.{md,mdx}", base: "src/content/posts" }),
  schema: z.object({
    title: z.string().optional(),
    meta_title: z.string().optional(),
    description: z.string().optional(),
    date: z.date().optional(),
    image: z.string().optional(),
    categories: z.array(z.string()).optional(),
    featured: z.boolean().optional(),
    draft: z.boolean().optional(),
  }),
});

// Export collections
export const collections = {
  about: aboutCollection,
  contact: contactCollection,
  homepage: homepageCollection,
  pages: pagesCollection,
  posts: postsCollection,
};
