import { createDirectus, rest, readItems } from "@directus/sdk";
import config from "@/config/config.json";
import type { Schema, Category, Post, LandingPage } from "./types";


class DirectusCms {

  directus;

  constructor() {
      this.directus = createDirectus<Schema>(config.site.directus_url).with(rest());
  }

  async GetAllCategories(): Promise<Category[]> {
    return await this.directus.request(readItems('categories',{
      filter:{
        "status": {
          _eq: 'published'
        },
      }
    }));
  }

  async GetAllPosts(): Promise<Post[]> {
    return await this.directus.request(readItems('posts',{
      filter:{
        "status": {
          _eq: 'published'
        },
      },
      limit: 20,
    }));
  }

  async GetPostsByCategory(category: string): Promise<Post[]> {
    return await this.directus.request(readItems('posts',{
      filter:{
        "status": {
          _eq: 'published'
        },
        "category":{
          category:{
              _eq: category
          }
      }},
      limit: 20,
    }));
  }

  async GetAllLandingPages(): Promise<LandingPage[]> {
    return await this.directus.request(readItems('landing_pages',{
      filter:{
        "status": {
          _eq: 'published'
        },
      }
    }));
  }


}

const Directus = new DirectusCms();
export default Directus;
