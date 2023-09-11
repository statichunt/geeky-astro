import { defineCollection, z } from "astro:content";

const metaObject = z.object({
  title: z.string(),
  meta_title: z.string().optional(),
  description: z.string().optional(),
  image: z.string().optional(),
  draft: z.boolean().optional(),
});

// Pages collection schema
const pagesCollection = defineCollection({
  schema: metaObject,
});

// Export collections
export const collections = {
  pages: pagesCollection,
};
