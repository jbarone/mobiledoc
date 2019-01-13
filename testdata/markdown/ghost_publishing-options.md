The Ghost editor has everything you need to fully optimise your content. This is where you can add tags and authors, feature a post, or turn a post into a page.

> Access the post settings menu in the top right hand corner of the editor.

## Post feature image

Insert your post feature image from the very top of the post settings menu. Consider resizing or optimising your image first to ensure it's an appropriate size.

## Structured data & SEO

Customise your social media sharing cards for Facebook and Twitter, enabling you to add custom images, titles and descriptions for social media.

There’s no need to hard code your meta data. You can set your meta title and description using the post settings tool, which has a handy character guide and SERP preview.

Ghost will automatically implement structured data for your publication using JSON-LD to further optimise your content.

```
{
    "@context": "https://schema.org",
    "@type": "Article",
    "publisher": {
        "@type": "Organization",
        "name": "Publishing options",
        "logo": "https://static.ghost.org/ghost-logo.svg"
    },
    "author": {
        "@type": "Person",
        "name": "Ghost",
        "url": "http://demo.ghost.io/author/ghost/",
        "sameAs": []
    },
    "headline": "Publishing options",
    "url": "http://demo.ghost.io/publishing-options",
    "datePublished": "2018-08-08T11:44:00.000Z",
    "dateModified": "2018-08-09T12:06:21.000Z",
    "keywords": "Getting Started",
    "description": "The Ghost editor has everything you need to fully optimise your content. This is where you can add tags and authors, feature a post, or turn a post into a page.",
    }
}
    
```

You can test that the structured data [schema](https://schema.org/) on your site is working as it should using [Google’s structured data tool](https://search.google.com/structured-data/testing-tool).

## Code Injection

This tool allows you to inject code on a per post or page basis, or across your entire site. This means you can modify CSS, add unique tracking codes, or add other scripts to the head or foot of your publication without making edits to your theme files.

**To add code site-wide**, use the code injection tool [in the main admin menu](/ghost/settings/code-injection/). This is useful for adding a Facebook Pixel, a Google Analytics tracking code, or to start tracking with any other analytics tool.

**To add code to a post or page**, use the code injection tool within the post settings menu. This is useful if you want to add art direction, scripts or styles that are only applicable to one post or page.

From here, you might be interested in managing some more specific [admin settings](/admin-settings/)!

