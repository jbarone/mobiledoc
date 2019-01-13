Ghost comes with a beautiful default theme called Casper, which is designed to be a clean, readable publication layout and can be adapted for most purposes. However, Ghost can also be completely themed to suit your needs. Rather than just giving you a few basic settings which act as a poor proxy for code, we just let you write code.

There are a huge range of both free and premium pre-built themes which you can get from the [Ghost Theme Marketplace](http://marketplace.ghost.org), or you can create your own from scratch.

{{< figure src="https://static.ghost.org/v1.0.0/images/marketplace.jpg" caption="Anyone can write a completely custom Ghost theme with some solid knowledge of HTML and CSS" >}}

Ghost themes are written with a templating language called handlebars, which has a set of dynamic helpers to insert your data into template files. For example: `{{author.name}}` outputs the name of the current author.

The best way to learn how to write your own Ghost theme is to have a look at [the source code for Casper](https://github.com/TryGhost/Casper), which is heavily commented and should give you a sense of how everything fits together.

* `default.hbs` is the main template file, all contexts will load inside this file unless specifically told to use a different template.
* `post.hbs` is the file used in the context of viewing a post.
* `index.hbs` is the file used in the context of viewing the home page.
* and so on

We've got [full and extensive theme documentation](https://themes.ghost.org/v2.0.0/docs) which outlines every template file, context and helper that you can use.

If you want to chat with other people making Ghost themes to get any advice or help, there's also a **themes** section on our [public Ghost forum](https://forum.ghost.org/c/themes).

