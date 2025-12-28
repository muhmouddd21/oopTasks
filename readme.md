🎯 Senior Design Task — Notification System
الفكرة العامة

هتبني نظام إشعارات (Notifications)
النظام ده يقدر يبعث إشعارات بطرق مختلفة:

Email

SMS

Push Notification
ومن غير ما نربط السيستم بطريقة واحدة.

الهدف الحقيقي:
Depend on behavior, not implementation

🧠 المطلوب منك (قبل أي كود)
اسأل نفسك:

مين المستهلك؟

إيه السلوك المشترك؟

إيه اللي هيتغيّر؟

فين الـ orchestration؟

فين الـ entities؟

📦 Requirements (من غير حلول)
1️⃣ Notification Behavior

أي Notification في السيستم لازم:

تبعت رسالة

يكون ليها recipient

يكون ليها cost (تكلفة إرسال)

❗ ممنوع:

if/else على النوع

switch على strings

2️⃣ Notification Types

اعمل أنواع مختلفة:

EmailNotification

SMSNotification

PushNotification

كل واحد:

له طريقة إرسال مختلفة

له تكلفة مختلفة

لكن السيستم ما يعرفش النوع

3️⃣ Notification Service (Orchestrator)

اعمل Service:

يستقبل مجموعة Notifications

يبعثهم

يحسب إجمالي التكلفة

📌 service ده:

ما يعرفش implementation

يعتمد على abstraction فقط

4️⃣ Rules مهمة جدًا

ممنوع أي type casting

ممنوع if/else حسب النوع

ممنوع service يعرف تفاصيل

5️⃣ Extension Test (مهم 🔥)

تخيّل:

بعد أسبوع قالوا عايزين WhatsAppNotification

اسأل نفسك:

هل هتعدّل Service؟

ولا هتضيف implementation بس؟

لو احتجت تعدّل Service → التصميم غلط ❌

🧩 Deliverables (إنت تعملهم)

من غير ما تحل هنا:

Interface واحدة أو أكتر (behavior)

Struct base (لو محتاج)

Implementations مختلفة

Service واحد

main بسيطة تختبر polymorphism

🧠 Self-Review Checklist (لازم تجاوب عليها)

بعد ما تخلص، اسأل نفسك:

 هل service يعرف أي نوع notification؟

 هل أقدر أضيف نوع جديد بدون تعديل كود قديم؟

 هل عندي if/else حسب النوع؟ (لو آه → ❌)

 هل الاتجاه معكوس صح؟

 هل أقدر أعمل FakeNotification للاختبار؟

💡 Hint (تفكير مش حل)

Notification ≠ Email

Notification = behavior

Service = orchestrator

Types = details

🏁 المطلوب منك دلوقتي

ابدأ بالترتيب ده:

1️⃣ اكتب interface فقط
2️⃣ اقفل الملف
3️⃣ اسأل نفسك: مين هيستهلك ده؟
4️⃣ بعدها بس اكتب implementation واحد
5️⃣ اختبر polymorphism
6️⃣ كمل الباقي

لما تخلص ⏭️

ابعتلي:

الـ interface

struct واحد

service skeleton

وأنا أعملك:

Senior Code Review

وأقولك:

فين التفكير الصح

فين أي smell

وإزاي تحسّن التصميم

🔥
إنت دلوقتي بتتعلّم تفكّر كمُهندس مش كمُنفّذ
يلا ابدأ، وأنا مستنيك.