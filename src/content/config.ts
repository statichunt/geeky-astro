import { defineCollection, z } from "astro:content";

const metaObject = z.object({
  title: z.string(),
  meta_title: z.string().optional(),
  description: z.string().optional(),
  image: z.string().optional(),
  draft: z.boolean().optional(),
});

// Post collection schema
const postCollection = defineCollection({
  schema: metaObject.merge(
    z.object({
      categories: z.string().array().optional(),
      featured: z.boolean().optional(),
      date: z.date().optional(),
      author: z.string().optional(),
    }),
  ),
});

// Author collection schema
const authorsCollection = defineCollection({
  schema: z.object({
    title: z.string(),
    meta_title: z.string().optional(),
    email: z.string().optional(),
    image: z.string().optional(),
    description: z.string().optional(),
    social: z
      .array(
        z
          .object({
            name: z.string().optional(),
            icon: z.string().optional(),
            link: z.string().optional(),
          })
          .optional(),
      )
      .optional(),
    draft: z.boolean().optional(),
  }),
});

// Pages collection schema
const pagesCollection = defineCollection({
  schema: metaObject,
});

// Export collections
export const collections = {
  posts: postCollection,
  authors: authorsCollection,
  pages: pagesCollection,
};
