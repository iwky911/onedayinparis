# -*- encoding:utf-8 -*-

# parses raw times and converts them to a graph; may require several reviews of the nodes
# to make sure that slight punctuation variations don't get set as different nodes
# data gathered manually from http://www.ratp.fr/fr/upload/docs/application/pdf/2011-11/premiers_derniers_departs.pdf

data = %{Line	Start	End	Duration	One Way?
1	Chateau de Vincennes			
1	Bérault	Château de Vincennes	0:01	
1	Saint-Mandé–Tourelle	Bérault	0:01	
1	Porte de Vincennes	Saint-Mandé–Tourelle	0:01	
1	Nation	Porte de Vincennes	0:02	
1	Reuilly–Diderot	Nation	0:02	
1	Gare de Lyon	Reuilly–Diderot	0:01	
1	Bastille	Gare de Lyon	0:02	
1	Saint-Paul	Bastille	0:02	
1	Hôtel de Ville	Saint-Paul	0:01	
1	Châtelet	Hôtel de Ville	0:02	
1	Louvre–Rivoli	Châtelet	0:01	
1	Palais Royal–Musée du Louvre 	Louvre–Rivoli	0:01	
1	Tuileries	Palais Royal–Musée du Louvre 	0:01	
1	Concorde	Tuileries	0:01	
1	Champs Élysées–Clemenceau 	Concorde	0:02	
1	Franklin D. Roosevelt	Champs Élysées–Clemenceau 	0:02	
1	George V	Franklin D. Roosevelt	0:01	
1	Charles de Gaulle–Étoile 	George V	0:01	
1	Argentine	Charles de Gaulle–Étoile 	0:02	
1	Porte Maillot	Argentine	0:01	
1	Les Sablons	Porte Maillot	0:02	
1	Pont de Neuilly	Les Sablons	0:01	
1	Esplanade de La Défense	Pont de Neuilly	0:02	
1	La Défense	Esplanade de La Défense	0:02	
2	Nation			
2	Avron	Nation	0:01	
2	Alexandre Dumas 	Avron	0:01	
2	Philippe Auguste	Alexandre Dumas 	0:01	
2	Père Lachaise 	Philippe Auguste	0:02	
2	Ménilmontant 	Père Lachaise 	0:01	
2	Couronnes	Ménilmontant 	0:01	
2	Belleville	Couronnes	0:01	
2	Colonel Fabien	Belleville	0:02	
2	Jaurès	Colonel Fabien	0:01	
2	Stalingrad	Jaurès	0:02	
2	La Chapelle 	Stalingrad	0:01	
2	Barbès-Rochechouart 	La Chapelle 	0:02	
2	Anvers	Barbès-Rochechouart 	0:01	
2	Pigalle	Anvers	0:02	
2	Blanche	Pigalle	0:01	
2	Place de Clichy	Blanche	0:02	
2	Rome	Place de Clichy	0:01	
2	Villiers	Rome	0:02	
2	Monceau	Villiers	0:01	
2	Courcelles	Monceau	0:01	
2	Ternes	Courcelles	0:02	
2	Charles de Gaulle–Étoile 	Ternes	0:01	
2	Victor Hugo	Charles de Gaulle–Étoile 	0:02	
2	Porte Dauphine	Victor Hugo	0:03	
3	Gallieni			
3	Porte de Bagnolet 	Gallieni	0:01	
3	Gambetta	Porte de Bagnolet 	0:01	
3	Père Lachaise 	Gambetta	0:02	
3	Rue Saint-Maur 	Père Lachaise 	0:01	
3	Parmentier 	Rue Saint-Maur 	0:01	
3	République 	Parmentier 	0:02	
3	Temple	République 	0:01	
3	Arts et Métiers 	Temple	0:01	
3	Réaumur-Sébastopol 	Arts et Métiers 	0:01	
3	Sentier	Réaumur-Sébastopol 	0:01	
3	Bourse 	Sentier	0:02	
3	Quatre-Septembre 	Bourse 	0:01	
3	Opéra 	Quatre-Septembre 	0:01	
3	Havre-Caumartin 	Opéra 	0:01	
3	Saint-Lazare	Havre-Caumartin 	0:01	
3	Europe	Saint-Lazare	0:02	
3	Villiers	Europe	0:01	
3	Malesherbes 	Villiers	0:02	
3	Wagram	Malesherbes 	0:01	
3	Pereire	Wagram	0:01	
3	Porte de Champerret 	Pereire	0:01	
3	Louise Michel 	Porte de Champerret 	0:01	
3	Anatole France	Louise Michel 	0:02	
3	Pont de Levallois	Anatole France	0:02	
3bis	Porte des Lilas			
3bis	Saint-Fargeau 	Porte des Lilas	0:01	
3bis	Pelleport 	Saint-Fargeau 	0:01	
3bis	Gambetta	Pelleport 	0:01	
4	Porte de Clignancourt			
4	Simplon	Porte de Clignancourt	0:01	
4	Marcadet–Poissonniers	Simplon	0:01	
4	Château Rouge 	Marcadet–Poissonniers	0:01	
4	Barbès–Rochechouart 	Château Rouge 	0:01	
4	Gare du Nord	Barbès–Rochechouart 	0:02	
4	Gare de l’Est	Gare du Nord	0:02	
4	Château d’Eau 	Gare de l’Est	0:01	
4	Strasbourg–Saint-Denis 	Château d’Eau 	0:01	
4	Réaumur–Sébastopol 	Strasbourg–Saint-Denis 	0:01	
4	Étienne Marcel	Réaumur–Sébastopol 	0:02	
4	Les Halles	Étienne Marcel	0:01	
4	Châtelet	Les Halles	0:01	
4	Cité	Châtelet	0:02	
4	Saint-Michel	Cité	0:01	
4	Odéon 	Saint-Michel	0:01	
4	Saint-Germain-des-Prés 	Odéon 	0:01	
4	Saint-Sulpice 	Saint-Germain-des-Prés 	0:01	
4	Saint-Placide 	Saint-Sulpice 	0:02	
4	Montparnasse–Bienvenüe 	Saint-Placide 	0:01	
4	Vavin	Montparnasse–Bienvenüe 	0:01	
4	Raspail 	Vavin	0:02	
4	Denfert-Rochereau 	Raspail 	0:01	
4	Mouton-Duvernet	Denfert-Rochereau 	0:01	
4	Alésia	Mouton-Duvernet	0:01	
4	Porte d’Orléans	Alésia	0:02	
5	Bobigny–Pablo Picasso			
5	Bobigny-Pantin–Raymond 	Bobigny–Pablo Picasso	0:03	
5	Queneau Église de Pantin	Bobigny-Pantin–Raymond 	0:02	
5	Hoche	Queneau Église de Pantin	0:01	
5	Porte de Pantin	Hoche	0:02	
5	Ourcq	Porte de Pantin	0:01	
5	Laumière	Ourcq	0:01	
5	Jaurès	Laumière	0:02	
5	Stalingrad	Jaurès	0:01	
5	Gare du Nord	Stalingrad	0:02	
5	Gare de l’Est 	Gare du Nord	0:03	
5	Jacques Bonsergent 	Gare de l’Est 	0:02	
5	République 	Jacques Bonsergent 	0:01	
5	Oberkampf 	République 	0:02	
5	Richard-Lenoir 	Oberkampf 	0:01	
5	Bréguet-Sabin 	Richard-Lenoir 	0:02	
5	Bastille	Bréguet-Sabin 	0:01	
5	Quai de la Rapée 	Bastille	0:02	
5	Gare d’Austerlitz 	Quai de la Rapée 	0:02	
5	Saint-Marcel 	Gare d’Austerlitz 	0:01	
5	Campo-Formio 	Saint-Marcel 	0:01	
5	Place d’Italie	Campo-Formio 	0:04	
6	Nation			
6	Picpus	Nation	0:01	
6	Bel Air	Picpus	0:01	
6	Daumesnil	Bel Air	0:01	
6	Dugommier	Daumesnil	0:01	
6	Bercy	Dugommier	0:01	
6	Quai de la Gare 	Bercy	0:02	
6	Chevaleret	Quai de la Gare 	0:01	
6	Nationale	Chevaleret	0:01	
6	Place d’Italie	Nationale	0:01	
6	Corvisart	Place d’Italie	0:01	
6	Glacière	Corvisart	0:02	
6	Saint-Jacques 	Glacière	0:01	
6	Denfert-Rochereau 	Saint-Jacques 	0:01	
6	Raspail	Denfert-Rochereau 	0:01	
6	Edgar Quinet 	Raspail	0:01	
6	Montparnasse–Bienvenüe 	Edgar Quinet 	0:01	
6	Pasteur	Montparnasse–Bienvenüe 	0:02	
6	Sèvres–Lecourbe 	Pasteur	0:01	
6	Cambronne	Sèvres–Lecourbe 	0:01	
6	La Motte-Picquet–Grenelle 	Cambronne	0:01	
6	Dupleix	La Motte-Picquet–Grenelle 	0:02	
6	Bir-Hakeim	Dupleix	0:01	
6	Passy	Bir-Hakeim	0:01	
6	Trocadéro	Passy	0:01	
6	Boissière	Trocadéro	0:01	
6	Kléber	Boissière	0:02	
6	Charles de Gaulle–Étoile	Kléber	0:02	
7	La Courneuve–8 mai 1945			
7	Fort d’Aubervilliers 	La Courneuve–8 mai 1945	0:01	
7	Aubervilliers–Pantin–Q. Chemins 	Fort d’Aubervilliers 	0:02	
7	Porte de la Villette 	Aubervilliers–Pantin–Q. Chemins 	0:02	
7	Corentin Cariou 	Porte de la Villette 	0:02	
7	Crimée 	Corentin Cariou 	0:01	
7	Riquet 	Crimée 	0:01	
7	Stalingrad 	Riquet 	0:01	
7	Louis Blanc 	Stalingrad 	0:02	
7	Château Landon 	Louis Blanc 	0:01	
7	Gare de l’Est 	Château Landon 	0:01	
7	Poissonnière 	Gare de l’Est 	0:02	
7	Cadet 	Poissonnière 	0:01	
7	Le Peletier 	Cadet 	0:01	
7	Chaussée d’Antin–La Fayette 	Le Peletier 	0:02	
7	Opéra 	Chaussée d’Antin–La Fayette 	0:01	
7	Pyramides 	Opéra 	0:02	
7	Palais-Royal–Musée du Louvre 	Pyramides 	0:01	
7	Pont Neuf 	Palais-Royal–Musée du Louvre 	0:02	
7	Châtelet 	Pont Neuf 	0:01	
7	Pont Marie 	Châtelet 	0:02	
7	Sully–Morland 	Pont Marie 	0:01	
7	Jussieu 	Sully–Morland 	0:02	
7	Place Monge 	Jussieu 	0:01	
7	Censier–Daubenton 	Place Monge 	0:01	
7	Les Gobelins 	Censier–Daubenton 	0:01	
7	Place d’Italie 	Les Gobelins 	0:02	
7	Tolbiac 	Place d’Italie 	0:01	
7	Maison Blanche	Tolbiac 	0:01	
7	Porte d’Italie	Maison Blanche	0:02	
7	Porte de Choisy	Porte d’Italie	0:01	
7	Porte d’Ivry	Porte de Choisy	0:01	
7	Pierre et Marie Curie	Porte d’Ivry	0:02	
7	Mairie d’ Ivry	Pierre et Marie Curie	0:03	
7	Maison Blanche			
7	Le Kremlin-Bicêtre	Maison Blanche	0:02	
7	Villejuif–Léo Lagrange	Le Kremlin-Bicêtre	0:01	
7	Villejuif–P. Vaillant-Couturier	Villejuif–Léo Lagrange	0:02	
7	Villejuif–Louis Aragon	Villejuif–P. Vaillant-Couturier	0:03	
7bis	Pré-Saint-Gervais			
7bis	Danube	Pré-Saint-Gervais	0:01	-1
7bis	Botzaris	Danube	0:01	-1
7bis	Buttes Chaumont	Botzaris	0:01	
7bis	Bolivar	Buttes Chaumont	0:01	
7bis	Jaurès	Bolivar	0:01	
7bis	Louis Blanc	Jaurès	0:03	
7bis	Botzaris			
7bis	Place des Fêtes	Botzaris	0:02	1
7bis	Pré-Saint-Gervais	Place des Fêtes	0:02	1
8	Balard			
8	Lourmel	Balard	0:01	
8	Boucicaut	Lourmel	0:01	
8	Félix Faure	Boucicaut	0:01	
8	Commerce	Félix Faure	0:01	
8	La Motte-Picquet–Grenelle	Commerce	0:01	
8	École Militaire	La Motte-Picquet–Grenelle	0:02	
8	La Tour-Maubourg	École Militaire	0:01	
8	Invalides	La Tour-Maubourg	0:01	
8	Concorde	Invalides	0:02	
8	Madeleine	Concorde	0:01	
8	Opéra	Madeleine	0:02	
8	Richelieu–Drouot	Opéra	0:01	
8	Grands Boulevards	Richelieu–Drouot	0:01	
8	Bonne Nouvelle	Grands Boulevards	0:01	
8	Strasbourg–Saint-Denis	Bonne Nouvelle	0:01	
8	République	Strasbourg–Saint-Denis	0:02	
8	Filles du Calvaire	République	0:01	
8	Saint-Sébastien–Froissart	Filles du Calvaire	0:01	
8	Chemin Vert	Saint-Sébastien–Froissart	0:01	
8	Bastille	Chemin Vert	0:01	
8	Ledru-Rollin	Bastille	0:02	
8	Faidherbe–Chaligny	Ledru-Rollin	0:01	
8	Reuilly–Diderot	Faidherbe–Chaligny	0:01	
8	Montgallet	Reuilly–Diderot	0:01	
8	Daumesnil	Montgallet	0:02	
8	Michel Bizot	Daumesnil	0:01	
8	Porte Dorée	Michel Bizot	0:01	
8	Porte de Charenton	Porte Dorée	0:02	
8	Liberté	Porte de Charenton	0:02	
8	Charenton-Écoles	Liberté	0:01	
8	École Vétérinaire de Maisons-Alfort 	Charenton-Écoles	0:02	
8	Maisons-Alfort–Stade	École Vétérinaire de Maisons-Alfort 	0:02	
8	Maisons-Alfort–Les Juilliottes 	Maisons-Alfort–Stade	0:02	
8	Créteil–L’Échat	Maisons-Alfort–Les Juilliottes 	0:01	
8	Créteil–Université	Créteil–L’Échat	0:02	
8	Créteil–Préfecture	Créteil–Université	0:02	
8	Pointe du Lac	Créteil–Préfecture	0:03	
9	Pont de Sèvres	Pointe du Lac		
9	Billancourt	Pont de Sèvres	0:01	
9	Marcel Sembat	Billancourt	0:01	
9	Porte de Saint-Cloud	Marcel Sembat	0:02	
9	Exelmans	Porte de Saint-Cloud	0:02	
9	Michel-Ange–Molitor	Exelmans	0:01	
9	Michel-Ange–Auteuil	Michel-Ange–Molitor	0:01	
9	Jasmin	Michel-Ange–Auteuil	0:02	
9	Ranelagh	Jasmin	0:01	
9	La Muette	Ranelagh	0:01	
9	Rue de la Pompe	La Muette	0:02	
9	Trocadéro	Rue de la Pompe	0:01	
9	Iéna	Trocadéro	0:01	
9	Alma–Marceau	Iéna	0:02	
9	Franklin D. Roosevelt	Alma–Marceau	0:02	
9	Saint-Philippe-du-Roule	Franklin D. Roosevelt	0:01	
9	Miromesnil	Saint-Philippe-du-Roule	0:01	
9	Saint-Augustin	Miromesnil	0:02	
9	Havre–Caumartin	Saint-Augustin	0:01	
9	Chaussée d’Antin–La Fayette	Havre–Caumartin	0:01	
9	Richelieu–Drouot	Chaussée d’Antin–La Fayette	0:01	
9	Grands Boulevards	Richelieu–Drouot	0:02	
9	Bonne Nouvelle	Grands Boulevards	0:01	
9	Strasbourg–Saint-Denis	Bonne Nouvelle	0:01	
9	République	Strasbourg–Saint-Denis	0:02	
9	Oberkampf	République	0:01	
9	Saint-Ambroise	Oberkampf	0:02	
9	Voltaire	Saint-Ambroise	0:01	
9	Charonne	Voltaire	0:02	
9	Rue des Boulets	Charonne	0:01	
9	Nation	Rue des Boulets	0:01	
9	Buzenval	Nation	0:02	
9	Maraîchers	Buzenval	0:01	
9	Porte de Montreuil	Maraîchers	0:01	
9	Robespierre	Porte de Montreuil	0:02	
9	Croix de Chavaux	Robespierre	0:02	
9	Mairie de Montreuil	Croix de Chavaux	0:03	
10	Boulogne–Pont de Saint-Cloud			
10	Boulogne–Jean Jaurès	Boulogne–Pont de Saint-Cloud	0:01	
10	Porte d’Auteuil	Boulogne–Jean Jaurès	0:03	1
10	Michel-Ange–Auteuil	Porte d’Auteuil	0:01	1
10	Église d’Auteuil	Michel-Ange–Auteuil	0:01	1
10	Javel–André Citroën	Église d’Auteuil	0:01	1
10	Charles Michels	Javel–André Citroën	0:02	
10	Avenue Émile Zola	Charles Michels	0:01	
10	La Motte-Picquet–Grenelle	Avenue Émile Zola	0:01	
10	Ségur	La Motte-Picquet–Grenelle	0:02	
10	Duroc	Ségur	0:02	
10	Vaneau	Duroc	0:01	
10	Sèvres–Babylone	Vaneau	0:01	
10	Mabillon	Sèvres–Babylone	0:02	
10	Odéon	Mabillon	0:01	
10	Cluny–La Sorbonne	Odéon	0:01	
10	Maubert–Mutualité	Cluny–La Sorbonne	0:01	
10	Cardinal Lemoine	Maubert–Mutualité	0:01	
10	Jussieu	Cardinal Lemoine	0:02	
10	Gare d’Austerlitz	Jussieu	0:03	
10	Boulogne–Jean Jaurès			
10	Michel-Ange–Molitor	Boulogne–Jean Jaurès	0:02	-1
10	Chardon–Lagache	Michel-Ange–Molitor	0:01	-1
10	Mirabeau	Chardon–Lagache	0:03	-1
10	Javel–André Citroën	Mirabeau	0:01	-1
11	Mairie des Lilas			
11	Porte des Lilas	Mairie des Lilas	0:01	
11	Télégraphe	Porte des Lilas	0:01	
11	Place des Fêtes	Télégraphe	0:01	
11	Jourdain	Place des Fêtes	0:01	
11	Pyrénées	Jourdain	0:01	
11	Belleville	Pyrénées	0:02	
11	Goncourt	Belleville	0:01	
11	République	Goncourt	0:01	
11	Arts et Métiers	République	0:02	
11	Rambuteau	Arts et Métiers	0:01	
11	Hôtel de Ville	Rambuteau	0:02	
11	Châtelet	Hôtel de Ville	0:02	
12	Mairie d’Issy			
12	Corentin Celton	Mairie d’Issy	0:01	
12	Porte de Versailles	Corentin Celton	0:01	
12	Convention	Porte de Versailles	0:02	
12	Vaugirard	Convention	0:01	
12	Volontaires	Vaugirard	0:02	
12	Pasteur	Volontaires	0:01	
12	Falguière	Pasteur	0:01	
12	Montparnasse–Bienvenüe	Falguière	0:01	
12	Notre-Dame-des-Champs	Montparnasse–Bienvenüe	0:02	
12	Rennes	Notre-Dame-des-Champs	0:01	
12	Sèvres–Babylone	Rennes	0:01	
12	Rue du Bac	Sèvres–Babylone	0:01	
12	Solférino	Rue du Bac	0:01	
12	Assemblée Nationale	Solférino	0:01	
12	Concorde	Assemblée Nationale	0:02	
12	Madeleine	Concorde	0:01	
12	Saint-Lazare	Madeleine	0:02	
12	Trinité–D’Estienne d’Orves	Saint-Lazare	0:01	
12	Notre-Dame-de-Lorette	Trinité–D’Estienne d’Orves	0:01	
12	Saint-Georges	Notre-Dame-de-Lorette	0:02	
12	Pigalle	Saint-Georges	0:01	
12	Abbesses	Pigalle	0:01	
12	Lamark–Caulaincourt	Abbesses	0:01	
12	Jules Joffrin	Lamark–Caulaincourt	0:02	
12	Marcadet–Poissonniers	Jules Joffrin	0:01	
12	Marx Dormoy	Marcadet–Poissonniers	0:02	
12	Porte de la Chapelle	Marx Dormoy	0:03	
13	Châtillon–Montrouge			
13	Malakoff–Rue Étienne Dolet	Châtillon–Montrouge	0:02	
13	Malakoff–Plateau de Vanves	Malakoff–Rue Étienne Dolet	0:02	
13	Porte de Vanves	Malakoff–Plateau de Vanves	0:01	
13	Plaisance	Porte de Vanves	0:02	
13	Pernety	Plaisance	0:01	
13	Gaîté	Pernety	0:02	
13	Montparnasse–Bienvenüe	Gaîté	0:01	
13	Duroc	Montparnasse–Bienvenüe	0:02	
13	Saint-François-Xavier	Duroc	0:01	
13	Varenne	Saint-François-Xavier	0:01	
13	Invalides	Varenne	0:01	
13	Champs Élysées–Clemenceau	Invalides	0:02	
13	Miromesnil	Champs Élysées–Clemenceau	0:02	
13	Saint-Lazare	Miromesnil	0:02	
13	Liège	Saint-Lazare	0:02	
13	Place de Clichy	Liège	0:02	
13	La Fourche	Place de Clichy	0:01	
13	Brochant	La Fourche	0:02	
13	Porte de Clichy	Brochant	0:02	
13	Mairie de Clichy	Porte de Clichy	0:01	
13	Gabriel Péri	Mairie de Clichy	0:03	
13	Les Agnettes	Gabriel Péri	0:02	
13	Asnières–Gennevilliers - Les Courtilles	Les Agnettes	0:03	
13	La Fourche			
13	Guy Môquet	La Fourche	0:02	
13	Porte de Saint-Ouen	Guy Môquet	0:01	
13	Garibaldi	Porte de Saint-Ouen	0:02	
13	Mairie de Saint-Ouen	Garibaldi	0:02	
13	Carrefour Pleyel	Mairie de Saint-Ouen	0:02	
13	Saint-Denis–Porte de Paris	Carrefour Pleyel	0:02	
13	Basilique de Saint-Denis	Saint-Denis–Porte de Paris	0:01	
13	Saint-Denis–Université	Basilique de Saint-Denis	0:03	
14	Saint-Lazare			
14	Madeleine	Saint-Lazare	0:01	
14	Pyramides	Madeleine	0:02	
14	Châtelet	Pyramides	0:02	
14	Gare de Lyon	Châtelet	0:03	
14	Bercy	Gare de Lyon	0:02	
14	Cour Saint-Émilion	Bercy	0:01	
14	Bibliothèque François Mitterrand	Cour Saint-Émilion	0:02	
14	Olympiades	Bibliothèque François Mitterrand	0:02	}

class Node
  attr_reader :n_id, :name
  @@node_max = 0
  @@nodeTable = {}
  def initialize(name)
    @n_id = @@node_max
    @@node_max += 1
    @name = name
  end
  def to_s
    "#{@n_id},#{@name}"
  end
  def self.n_key(name)
    name.gsub(/[^a-zA-Z0-9]/,"")
  end
  def self.find_or_add(name)
    clean_name = name.strip
    n = @@nodeTable[n_key(clean_name)]
    if n.nil?
      n = Node.new(clean_name)
      @@nodeTable[n_key(clean_name)] = n
    end
    n
  end
  def self.find(name)
    return @@nodeTable[n_key(name)]
  end
  def self.all
    @@nodeTable.values
  end
end

class Edge
  attr_reader :e_id, :depart_id, :arrival_id, :duration, :line
  @@edge_max = 0
  @@edges = []
  def initialize(depart_id, arrival_id, duration, line)
    @e_id = @@edge_max; @@edge_max += 1
    @depart_id= depart_id; @arrival_id = arrival_id
    @duration = duration; @line = line
  end
  def self.add(e)
    @@edges << e
  end
  def self.all
    @@edges
  end
  def to_s
    "#{@depart_id},#{@arrival_id},#{@duration},#{@line}"
  end
end

def add(*args)
  return if args.length < 4
  line = args[0]
  duration = args[3].split(":")[1].to_i
  dep_node = Node.find_or_add args[1]
  arr_node = Node.find_or_add args[2]
  Edge.add(Edge.new dep_node.n_id, arr_node.n_id, duration, line) if args.length == 4 || args[4] == '1'
  Edge.add(Edge.new arr_node.n_id, dep_node.n_id, duration, line) if args.length == 4 || args[4] == '-1'
end

n_out = File.open("nodes.csv", "w")
e_out = File.open("edges.csv", "w")
data.split("\n").each_with_index { |l, i| add(*l.split("\t")) if i > 0 }
Node.all.sort_by(&:n_id).each {|n| n_out.print "#{n}\n"} #.sort_by(&:name)
Edge.all.each {|e| e_out.print "#{e}\n"}
