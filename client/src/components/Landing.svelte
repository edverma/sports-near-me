<script>
	import {  Heading, P, Button, Input, Span, Footer, FooterCopyright, FooterLinkGroup, FooterLink } from 'flowbite-svelte'
</script>

<div class="grid grid-cols-1 sm:grid-cols-2">
	<div class="h-screen mx-6">
		<div class="mt-6 flex">
			<Heading><Span class="bg-gradient-to-r from-violet-500 via-pink-500 to-pink-300" gradient>Friendmail</Span></Heading>
		</div>
		<div class="mt-[40%]">
			<Heading tag="h1" customSize="text-5xl lg:text-6xl font-bold">Send Mail to</Heading>
			<Heading tag="h1" customSize="text-5xl lg:text-6xl font-bold" color="text-pink-500">Your Friends.</Heading>
			<P size="xl" class="lg:text-2xl mt-12">Send a professionally designed newsletter to stay in touch and connect with friends across the country.</P>
		</div>
	</div>
	<div class="bg-black h-screen">
		<img class="my-[30%] md:my-[20%] mx-auto" src="/friendmail_logo.png"/>
	</div>
	<div class="sm:col-span-2">
		<Footer>
			<FooterCopyright href="/" by="Friendmail™" year={2023} />
			<!--
			<FooterLinkGroup ulClass="flex flex-wrap items-center mt-3 text-sm text-gray-500 dark:text-gray-400 sm:mt-0">
				<FooterLink href="/">About</FooterLink>
				<FooterLink href="/">Privacy Policy</FooterLink>
				<FooterLink href="/">Licensing</FooterLink>
				<FooterLink href="/">Contact</FooterLink>
			</FooterLinkGroup>
			-->
		</Footer>
	</div>
</div>
