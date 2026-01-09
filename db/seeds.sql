begin;

truncate posts, comments;

insert into posts (id, title, slug, body)
values (
  1,
  'Testing ASP.NET Core Applications Against a Real Database',
  'testing-asp-net-core-applications-against-database',
  'I am writing this post to record the steps necessary to set up integration testing in my side project, which I am building for my YouTube channel, [Make Programming Fun Again](https://www.youtube.com/@KarolMoroz).
The [project in question](https://github.com/moroz/FullStackAsp.Net-Courses) is a full stack application, built as an [ASP.NET Core](https://dotnet.microsoft.com/en-us/apps/aspnet) application.
It uses [PostgreSQL](https://www.postgresql.org/) and [Entity Framework Core](https://learn.microsoft.com/en-us/ef/core/) for data persistence, and exposes a [gRPC](https://grpc.io/)-based API.
On top of this API, a Web interface will be built using [SvelteKit](https://svelte.dev/docs/kit/introduction).

The code snippets in this post are largely inspired by Microsoft (2025)[^1], the recommendations of some friends, and a bunch of conversations with various LLMs.

If you want to follow along this walkthrough, clone the GitHub repository [moroz/FullStackAsp.Net-Courses](https://github.com/moroz/FullStackAsp.Net-Courses) and check out the branch `testing-tutorial` (starting at tag [2025-09-24](https://github.com/moroz/FullStackAsp.Net-Courses/tree/2025-09-24)):

```shell
git clone git@github.com:moroz/FullStackAsp.Net-Courses.git -b testing-tutorial
```

Disclaimer: This walkthrough is meant for use with a specific project. You may find it useful in other project, but it is not meant as a generic reference or tutorial.',
);

commit;
