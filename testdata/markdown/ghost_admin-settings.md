There are a couple of things to do next while you're getting set up:

# Make your site private

If you've got a publication that you don't want the world to see yet because it's not ready to launch, you can hide your Ghost site behind a basic shared pass-phrase.

You can toggle this preference on at the bottom of Ghost's [General Settings](/ghost/settings/general/):

{{< figure src="https://static.ghost.org/v1.0.0/images/private.png" >}}

Ghost will give you a short, randomly generated pass-phrase which you can share with anyone who needs access to the site while you're working on it. While this setting is enabled, all search engine optimisation features will be switched off to help keep your site under the radar.

Do remember though, this is _not_ secure authentication. You shouldn't rely on this feature for protecting important private data. It's just a simple, shared pass-phrase for some very basic privacy.

---

# Invite your team

Ghost has a number of different user roles for your team:

**Contributors**This is the base user level in Ghost. Contributors can create and edit their own draft posts, but they are unable to edit drafts of others or publish posts. Contributors are **untrusted** users with the most basic access to your publication.

**Authors**Authors are the 2nd user level in Ghost. Authors can write, edit  and publish their own posts. Authors are **trusted** users. If you don't trust users to be allowed to publish their own posts, they should be set as Contributors.

**Editors**Editors are the 3rd user level in Ghost. Editors can do everything that an Author can do, but they can also edit and publish the posts of others - as well as their own. Editors can also invite new Contributors+Authors to the site.

**Administrators**The top user level in Ghost is Administrator. Again, administrators can do everything that Authors and Editors can do, but they can also edit all site settings and data, not just content. Additionally, administrators have full access to invite, manage or remove any other user of the site.**The Owner**There is only ever one owner of a Ghost site. The owner is a special user which has all the same permissions as an Administrator, but with two exceptions: The Owner can never be deleted. And in some circumstances the owner will have access to additional special settings if applicable. For example: billing details, if using [**Ghost(Pro)**](https://ghost.org/pricing/).

> _It's a good idea to ask all of your users to fill out their user profiles, including bio and social links. These will populate rich structured data for posts and generally create more opportunities for themes to fully populate their design._

Next up: [Organising your content](/organising-content/) 

