begin;

delete from posts;
delete from comments;

insert into posts (id, title, slug, published_at, created_at, body)
values (
  '019ba31e-200e-717b-8a87-d540701485f5',
  'Testing ASP.NET Core Applications Against a Real Database',
  'testing-asp-net-core-applications-against-database',
  strftime('%s', '2025-09-29 12:00:00'),
  strftime('%s', '2025-09-29 12:00:00'),
  '<strong>This is an abridged version of a post originally published on [moroz.dev](https://moroz.dev/blog/testing-asp-net-core-applications-against-database/).</strong>

  I am writing this post to record the steps necessary to set up integration testing in my side project, which I am building for my YouTube channel, [Make Programming Fun Again](https://www.youtube.com/@KarolMoroz).
The [project in question](https://github.com/moroz/FullStackAsp.Net-Courses) is a full stack application, built as an [ASP.NET Core](https://dotnet.microsoft.com/en-us/apps/aspnet) application.
It uses [PostgreSQL](https://www.postgresql.org/) and [Entity Framework Core](https://learn.microsoft.com/en-us/ef/core/) for data persistence, and exposes a [gRPC](https://grpc.io/)-based API.
On top of this API, a Web interface will be built using [SvelteKit](https://svelte.dev/docs/kit/introduction).

The code snippets in this post are largely inspired by Microsoft (2025)[^1], the recommendations of some friends, and a bunch of conversations with various LLMs.

If you want to follow along this walkthrough, clone the GitHub repository [moroz/FullStackAsp.Net-Courses](https://github.com/moroz/FullStackAsp.Net-Courses) and check out the branch `testing-tutorial` (starting at tag [2025-09-24](https://github.com/moroz/FullStackAsp.Net-Courses/tree/2025-09-24)):

```shell
git clone git@github.com:moroz/FullStackAsp.Net-Courses.git -b testing-tutorial
```

Disclaimer: This walkthrough is meant for use with a specific project. You may find it useful in other project, but it is not meant as a generic reference or tutorial.

[^1]: Microsoft. (2025, March 25). *Integration tests in ASP.NET Core*. Microsoft Learn. https://learn.microsoft.com/en-us/aspnet/core/test/integration-tests?view=aspnetcore-9.0&pivots=xunit ([Web Archive](https://web.archive.org/web/20250927142022/https://learn.microsoft.com/en-us/aspnet/core/test/integration-tests?view=aspnetcore-9.0&pivots=xunit))'
);

insert into comments (id, post_id, signature, website, body)
values (
    '019ba389-d6cd-709f-af05-4eafa1f4977e',
    '019ba31e-200e-717b-8a87-d540701485f5',
    'Farmacia Hombres',
    'https://farmacia-hombres.com/potencia/cialis-generico.html',
    'Las ventas de productos online ofrecen distintas ventajas, pues cada vez es más seguro comprar por medio de la web, y esto permite a los productos que necesitamos de manera online. En este caso hablamos de una farmacia online barata en España para comprar Viagra, <a href="https://farmacia-hombres.com/potencia/cialis-generico.html" style="color:inherit;text-decoration:inherit; font-weight:inherit">cialis generico</a> y Levitra, cercana y con los productos que necesitas. Las ventas de productos online ofrecen distintas ventajas, pues cada vez es más seguro comprar por medio de la industria farmacéutica ha generado un nuevo mercado online en España. La entrega es una de los fármacos en las farmacias ordinarias'
),
(
  '019ba395-1878-743b-b540-e63a6a3fcaac',
  '019ba31e-200e-717b-8a87-d540701485f5',
  'خدمة الدعم',
 null,
  'شركة تنظيف بالرياض شركة تنظيف فلل بالرياض شركة تنظيف منازل بالرياض شركة تنظيف موكيت بالرياض شركة تنظيف مجالس بالرياض شركة تنظيف خزانات بالرياض شركة عزل خزانات بالرياض شركة عزل اسطح بالرياض شركة كشف تسربات المياه بالرياض شركة تسليك مجارى بالرياض شركة مكافحة حشرات بالرياض شركة رش مبيدات بالرياض شركة نقل اثاث بالرياض شركة تخزين اثاث بالرياض شركة تنظيف بيارات بالرياض شركة جلى بلاط بالرياض شركة تنظيف فلل بمكة شركات تنظيف منازل بجدة شركة نقل اثاث بجدة شركة تنظيف خزانات بجدة شركة كشف تسربات بجدة شركة تنظيف فلل بجدة شركة تنظيف موكيت بجدة شركة تسليك مجارى بجدة شركة عزل خزانات بجدة شركة تنظيف بيارات بجدة شركة تخزين عفش بجدة شركة تنظيف خزانات بمكة شركة تنظيف منازل بمكة شركة نقل اثاث بالدمام شركة عزل اسطح بالدمام شركة تنظيف خزانات بالدمام شركة مكافحة حشرات بالدمام شركة رش مبيدات بالدمام شركه تنظيف موكيت بالدمام شركة كشف تسربات المياه بالدمام شركه تنظيف فلل بالدمام شركة تنظيف بالجبيل شركة تسليك مجارى بالدمام شركة مكافحة حشرات بالجبيل شركة تنظيف منازل بالاحساء شركة مكافحة القوارض شركة مكافحة حشرات بالرياض شركة رش مبيدات بالرياض كشف تسربات المياه شركة مكافحة النمل الابيض بالرياض القضاء على البق شركة مكافحة الصراصير بالرياض مكافحة الفئران مكافحة حشرات الفراش شركة ابادة الحشرات شركة تخزين اثاث بجدة شركة نقل اثاث بجدة شركة رش مبيدات بجدة شركة تسليك مجارى بجدة شركة كشف تسرب المياه بجده شركات مكافحة الحشرات فى جدة شركة تنظيف المنازل فى جدة شركة تنظيف خزانات بجدة شركة تنظيف بيارات بجدة شركات مكافحة الحشرات في جدة شركات تنظيف المنازل في جدة شركة نقل عفش جدة شركة تنظيف خزانات بجدة شركة كشف تسرب المياه بجده شركة تنظيف فلل بجدة شركة تنظيف موكيت بجدة شركة تنظيف مجالس بجدة شركة تنظيف مسابح بجدة شركة تسليك مجارى بجدة شركة عزل مائى بجدة شركة تخزين عفش بجدة شركة نقل الاثاث فى مصر شركة نقل اثاث بالمنصورة شركة نقل اثاث بالاسكندرية شركة نقل اثاث بالمعادى شركة نقل اثاث فى مدينة نصر شركة نقل اثاث بمدينتى شركة نقل اثاث بالتجمع شركات نقل الاثاث بالرحاب شركة نقل اثاث بالقاهرة افضل شركة شحن فى مصر شركة نقل الاثاث فى مصر شركة نقل اثاث بالمنصورة شركة نقل اثاث بالاسكندرية شركة نقل اثاث بالمعادى شركة نقل اثاث فى مدينة نصر شركة نقل اثاث بمدينتى شركة نقل اثاث بالتجمع شركات نقل الاثاث بالرحاب شركة نقل اثاث بالقاهرة افضل شركة شحن فى مصر نقل عفش مصر شركة نقل اثاث بالمنصورة شركة نقل اثاث بالاسكندرية شركة نقل اثاث بالمعادى شركة نقل اثاث فى مدينة نصر شركة نقل اثاث بمدينتى شركة نقل اثاث بالتجمع شركات نقل الاثاث بالرحاب شركة نقل اثاث بالقاهرة افضل شركة شحن فى مصر شركة مكافحة حشرات ورش مبيدات برماح شركة تنظيف بيارات بشقراء شركة نقل أثاث بضرماء شركة تنظيف بضرماء شركة تخزين أثاث برماح شركة كشف تسربات المياه بالدلم شركة عزل أسطح بالمزاحمية شرك مكافحة حشرات ورش مبيدات بشقراء شركة مكافحة حشرات ورش مبيدات بالمجمعة شركة تنظيف فلل بالمجمعة شركة تنظيف بيارات بضرماء شركة كشف تسربات المياه برماح شركة كشف تسربات المياه بالمجمعة شركة تسليك مجاري بضرماء شركة تسليك مجاري بالدلم شركة عزل أسطح بشقراء شركة عزل أسطح بضرما شركة تنظيف بشقراء شركة عزل خزانات بضرما شركة تنظيف بيارات بضرما شركة تسليك مجاري برماح شركة نقل أثاث برماح شركة عزل خزانات بشقراء شركة تنظيف بيارات برماح شركة عزل أسطح بالدلم شركة نقل أثاث بشقراء شركة تخزين أثاث بشقراء شركة كشف تسربات المياه برأس التنورة شركة عزل خزانات برماح شركة تنظيف بالمزاحمية'
),
(
  '019ba394-e85f-71ab-8708-aa48c2d37961',
  '019ba31e-200e-717b-8a87-d540701485f5',
  'Jean Dupont',
  'http://www.frmontrepascher.eu',
  'Ce qui est pensé à cette http://www.frmontrepascher.eu montre exceptionnellement peu commune, en tout cas, c''est que c''était le principal chronographe Rolex dans un cas Oyster. Cette pièce, publiée en 1937, est la première à être équipée du calibre interne 10 1/2 ", un développement physiquement torsadé avec une capacité de chronographe à réflexion solitaire. Le Zerographe, qui a un boîtier de 32 mm et un brosse de poignet Oyster boulonné, en parle en premier à Rolex: une lunette tournante, qui se transformerait en une colonne vertébrale de nombreuses montres http://repliquemontreluxevideos.tumblr.com fausses et sportives Rolex. Aucune très grande partie de ces références n''a été publiée à la clôture.'
),
(
  '019ba397-2a1d-7113-8b9e-0063da7da979',
  '019ba31e-200e-717b-8a87-d540701485f5',
  'Replicas Relojes Marca',
  'http://www.replicasrelojesmarca.es',
  'Incorpora orejetas http://www.replicasrelojesmarca.es mejoradas y un bisel con cinco orificios de tornillo distintivos. La correa de caucho de este reloj resistente a los golpes se enrolla cómodamente en la muñeca para un ajuste correcto. Veamos el cuadrante, hay tres colores para elegir. Se completa con dos números arábigos sobredimensionados e índices en forma de barra, incrustados con compuesto luminoso. La pantalla indica http://www.relojesreplicasspain.com horas, minutos, segundos de pirateo y la fecha, presentada a las 3 en punto.'
),
(
  '019ba398-67b0-750b-8b38-db3735a2a81f',
  '019ba31e-200e-717b-8a87-d540701485f5',
  'Replicas Relojes Marca',
  'http://www.buenoreplicasrelojes.eu',
  'Si este calendario mantiene una operación normal, el ajuste manual solo una vez al año cuando en marzo desde febrero ingrese la fecha. El dispositivo de calendario http://www.buenoreplicasrelojes.eu puede mostrar el día, la fecha y el mes, y controla la visualización de día / noche. Nueva Ref. El dispositivo de sincronización 5905P también utiliza la última tecnología. Esto significa adherirse a la tradición, seguir usando el inicio cronógrafo controlado con rueda de columnas, detener y reiniciar. La diferencia es esa abrazadera de disco de embrague vertical controlada por rueda de http://www.mejorrelojesimitacion.eu columna, en lugar del nivel de las horquillas de rueda de embrague.'
);


commit;
